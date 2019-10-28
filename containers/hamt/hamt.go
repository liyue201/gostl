package hamt

import (
	"hash/fnv"
	"math/bits"
)

const (
	BITMAP_NODE = 0
	KV_NODE     = 1
	Fanout      = 6 //each bitmap node has 6 bits, so the max depth of tree is 64/6 = 10.666 = 11
	Mask        = (1 << Fanout) - 1
)

type Key []byte

type Entry interface {
	Type() int
	BitPos() uint64
}

type BitmapNode struct {
	bitmap   uint64
	children []Entry
	bitPos   uint64
}

type KvPair struct {
	key   []byte
	value interface{}
	next  *KvPair
}

type KvNode struct {
	hash   uint64
	kvList *KvPair
	bitPos uint64
}

type Hamt struct {
	root BitmapNode
}

func (this *BitmapNode) Type() int {
	return BITMAP_NODE
}

func (this *BitmapNode) BitPos() uint64 {
	return this.bitPos
}

func (this *BitmapNode) Index(pos uint64) int {
	return bits.OnesCount64((pos - 1) & this.bitmap)
}

func (this *BitmapNode) insert(depth int, hash uint64, kv *KvPair) {
	nBits := (hash >> (uint64(depth) * Fanout)) & Mask //hash in current node's bits
	bitPos := uint64(1) << nBits                       //hash in current bitmap's position
	if bitPos&this.bitmap == 0 {
		this.bitmap |= bitPos
		newChildren := make([]Entry, len(this.children)+1)
		kvNode := &KvNode{
			hash:   hash,
			kvList: kv,
			bitPos: bitPos,
		}
		index := this.Index(bitPos)
		newChildren[index] = kvNode
		for _, entry := range this.children {
			index = this.Index(entry.BitPos())
			newChildren[index] = entry
		}
		this.children = newChildren
	} else {
		index := this.Index(bitPos)
		entry := this.children[index]
		if entry.Type() == KV_NODE {
			kvNode := entry.(*KvNode)
			if kvNode.hash == hash {
				for iter := kvNode.kvList; iter != nil; iter = iter.next {
					if string(iter.key) == string(kv.key) {
						iter.value = kv.value
						return
					}
				}
				kv.next = kvNode.kvList
				kvNode.kvList = kv
			} else {
				bitmapNode := &BitmapNode{
					bitPos: bitPos,
				}
				bitmapNode.insert(depth+1, kvNode.hash, kvNode.kvList)
				bitmapNode.insert(depth+1, hash, kv)
				this.children[index] = bitmapNode
			}
		} else {
			entry.(*BitmapNode).insert(depth+1, hash, kv)
		}
	}
}

func (this *BitmapNode) find(depth int, hash uint64, key Key) interface{} {
	nBits := (hash >> (uint64(depth) * Fanout)) & Mask //hash in current node's bits
	bitPos := uint64(1) << nBits                       //hash in current bitmap's position
	if bitPos&this.bitmap == 0 {
		return nil
	}
	index := this.Index(bitPos)
	entry := this.children[index]
	if entry.Type() == KV_NODE {
		kvNode := entry.(*KvNode)
		if kvNode.hash != hash {
			return nil
		}

		for iter := kvNode.kvList; iter != nil; iter = iter.next {
			if string(iter.key) == string(key) {
				return iter.value
			}
		}
	} else {
		return entry.(*BitmapNode).find(depth+1, hash, key)
	}
	return nil
}

func (this *BitmapNode) traversal(f func(node *KvNode)) {
	for _, entry := range this.children {
		if entry.Type() == BITMAP_NODE {
			entry.(*BitmapNode).traversal(f)
		} else {
			f(entry.(*KvNode))
		}
	}
}

func (this *BitmapNode) erase(depth int, hash uint64, key Key) bool {
	nBits := (hash >> (uint64(depth) * Fanout)) & Mask //hash in current node's bits
	bitPos := uint64(1) << nBits                       //hash in current bitmap's position
	if bitPos&this.bitmap == 0 {
		return false
	}
	index := this.Index(bitPos)
	entry := this.children[index]
	if entry.Type() == KV_NODE {
		kvNode := entry.(*KvNode)
		if kvNode.hash != hash {
			return false
		}
		iter := kvNode.kvList
		var preIter *KvPair
		found := false
		for ; iter != nil; iter = iter.next {
			if string(iter.key) == string(key) {
				found = true
				break
			}
			preIter = iter
		}
		if found {
			// remove iter
			if preIter != nil {
				preIter.next = iter.next
			} else {
				kvNode.kvList = iter.next
			}
			if kvNode.kvList == nil {
				this.children[index] = nil
				this.bitmap &= ^bitPos
				newChildren := make([]Entry, len(this.children)-1)
				for _, entry := range this.children {
					if entry != nil {
						newIndex := this.Index(entry.BitPos())
						newChildren[newIndex] = entry
					}
				}
				this.children = newChildren
			}
			return true
		}
		return false
	} else {
		bitmapNode := entry.(*BitmapNode)
		ok := bitmapNode.erase(depth+1, hash, key)
		// change bitmapNode to kvNode, if the a bitmapNode has only one kvNode
		if ok && len(bitmapNode.children) == 1 && bitmapNode.children[0].Type() == KV_NODE {
			child := bitmapNode.children[0].(*KvNode)
			child.bitPos = bitmapNode.bitPos
			this.children[index] = child
		}
		return ok
	}
}

func (this *KvNode) Type() int {
	return KV_NODE
}

func (this *KvNode) BitPos() uint64 {
	return this.bitPos
}

// New new a Hamt(hash array map tree) instance
func New() *Hamt {
	return &Hamt{}
}

// Insert insert a key-value pair into hamt
func (this *Hamt) Insert(key Key, value interface{}) {
	keyHash := hash(key)
	this.root.insert(0, keyHash, &KvPair{key: key, value: value})
}

// Insert returns the value by the passed key, or nil if not found
func (this *Hamt) Get(key Key) interface{} {
	keyHash := hash(key)
	return this.root.find(0, keyHash, key)
}

// Insert erase the key-value pair in hamt, and returns true if succeed.
func (this *Hamt) Erase(key Key) bool {
	keyHash := hash(key)
	return this.root.erase(0, keyHash, key)
}

// Keys returns the keys in Hamt
func (this *Hamt) Keys() []Key {
	keys := make([]Key, 0)
	this.root.traversal(func(node *KvNode) {
		for kv := node.kvList; kv != nil; kv = kv.next {
			keys = append(keys, kv.key)
		}
	})
	return keys
}

// StringKeys returns the keys in Hamt
func (this *Hamt) StringKeys() []string {
	keys := make([]string, 0)
	this.root.traversal(func(node *KvNode) {
		for kv := node.kvList; kv != nil; kv = kv.next {
			keys = append(keys, string(kv.key))
		}
	})
	return keys
}

func hash(a []byte) uint64 {
	h := fnv.New64()
	h.Write(a)
	return h.Sum64()
}