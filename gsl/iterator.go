package gsl

type BidirectIterator interface{
	Next() bool

	Prev() bool

	Value() interface{}
	
	Index() int

	First() bool

	Assign(v interface{})

	Equal(t BidirectIterator) bool
}

type RandomAccessIterator interface {
	BidirectIterator
}

func Distance(start BidirectIterator, end BidirectIterator) int {
	d := end.Index() - start.Index()
	return d
}