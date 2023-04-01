package pair

type Pair struct {
	Front any
	Back  any
}

func MakePair(front any, back any) *Pair {
	return &Pair{Front: front, Back: back}
}

func (P *Pair) New(front any, back any) {
	P.Front = front
	P.Back = back
}

func (P *Pair) Equal(pair2 Pair) bool {
	return P.Front == pair2.Front && P.Back == pair2.Back
}

func (P *Pair) Fronts() any {
	return P.Front
}

func (P *Pair) Backs() any {
	return P.Back
}
