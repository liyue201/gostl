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
	BitPos(depth int) uint64
}

type BitmapNode struct {
	bitmap   uint64
	children []Entry
	pos      uint8 //position in parent array, in range [0, 64)
}

type KvPair struct {
	key   []byte
	value interface{}
	next  *KvPair
}

type KvNode struct {
	hash   uint64
	kvList *KvPair
}

type Hamt struct {
	root BitmapNode
}

func (this *BitmapNode) Type() int {
	return BITMAP_NODE
}

func (this *BitmapNode) BitPos(depth int) uint64 {
	return uint64(1) << this.pos
}

func (this *BitmapNode) Index(bitPos uint64) int {
	return bits.OnesCount64((bitPos - 1) & this.bitmap)
}

func (this *BitmapNode) insert(depth int, hash uint64, kv *KvPair) {
	pos := pos(hash, depth) //hash in current node's position
	bitPos := bitPos(pos)   //hash in current bitmap's position in bit
	if bitPos&this.bitmap == 0 {
		this.bitmap |= bitPos
		newChildren := make([]Entry, len(this.children)+1)
		kvNode := &KvNode{
			hash:   hash,
			kvList: kv,
		}
		index := this.Index(bitPos)
		newChildren[index] = kvNode
		for _, entry := range this.children {
			index = this.Index(entry.BitPos(depth))
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
					pos: pos,
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
	pos := pos(hash, depth) //hash in current node's position
	bitPos := bitPos(pos)   //hash in current bitmap's position in bit
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

type Visitor func(key, value interface{}) bool

func (this *BitmapNode) traversal(visitor Visitor) {
	for _, entry := range this.children {
		if entry.Type() == BITMAP_NODE {
			entry.(*BitmapNode).traversal(visitor)
		} else {
			node := entry.(*KvNode)
			for kv := node.kvList; kv != nil; kv = kv.next {
				if !visitor(kv.key, kv.value) {
					return
				}
			}
		}
	}
}

func (this *BitmapNode) erase(depth int, hash uint64, key Key) bool {
	pos := pos(hash, depth) //hash in current node's position
	bitPos := bitPos(pos)   //hash in current bitmap's position in bit
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
						newIndex := this.Index(entry.BitPos(depth))
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
			this.children[index] = child
		}
		return ok
	}
}

func (this *KvNode) Type() int {
	return KV_NODE
}

func (this *KvNode) BitPos(depth int) uint64 {
	return uint64(1) << pos(this.hash, depth)
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
	this.root.traversal(func(key, value interface{}) bool {
		keys = append(keys, key.(Key))
		return true
	})
	return keys
}

// StringKeys returns the keys in Hamt
func (this *Hamt) StringKeys() []string {
	keys := make([]string, 0)
	this.root.traversal(func(key, value interface{}) bool {
		keys = append(keys, string(key.(Key)))
		return true
	})
	return keys
}

// Traversal traversals elements in Hamt, it will stop until to the end or visitor returns false
func (this *Hamt) Traversal(visitor Visitor) {
	this.root.traversal(visitor)
}

func hash(a []byte) uint64 {
	h := fnv.New64()
	h.Write(a)
	return h.Sum64()
}

func pos(hash uint64, depth int) uint8 {
	return uint8((hash >> (uint64(depth) * Fanout)) & Mask)
}

func bitPos(pos uint8) uint64 {
	return uint64(1) << pos
}
