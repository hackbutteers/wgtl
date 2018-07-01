package gsl

type ListInterface interface {
	
	PushFront(v interface{})

	PushFrontv(v... interface{})

	PushBackv(v...interface{})

	Front() interface{}
	
	PopFront() interface{}
	
	PushBack(v interface{})
	
	PopBack() interface{}
	
	back() interface{}

	Get(index int) (interface{}, bool)

	Swap(index1 int, index2 int) bool

	FindIndex(index int) *ListIterator

	FindValue(v interface{}) *ListIterator

	RemoveIndex(index int)
	
	RemoveItr(it *ListIterator)
	RemoveValue(v interface{}, comp Comparator)

	Sort(comparator Comparator)
	
	Insert(index int, values ...interface{})

	Slice()[]interface{}

	Begin() *ListIterator

	End()   *ListIterator
	
	Rbegin() *ListReverseIterator
	
	Rend() *ListReverseIterator

	Add(v...interface{})
	
	getInternal(index int) *Lnode

	ContainerInterface

}

type ListDefaultAllocator struct {}


type ListIterator struct {
	index  int
	node   *Lnode
	list   *List
}

type ListReverseIterator struct {
	index  int
	node   *Lnode
	list   *List
}

type Lnode struct {
	next *Lnode
	prev *Lnode
	value interface{}
}

type List struct {
	alloc Allocator
	head *Lnode
	tail *Lnode
	size  int
}

func NewList() *List {
	al := &ListDefaultAllocator{}
	h := al.Allocate(1).(*Lnode)
	t := al.Allocate(1).(*Lnode)
	h.next = t
	h.prev = nil
	h.value  = nil
	t.next = nil
	t.value = nil
	t.prev = h
	return &List{al, h, t, 0}
}

func NewAllocList(al Allocator) *List {
	h := al.Allocate(1).(*Lnode)
	t := al.Allocate(1).(*Lnode)
	al.Construct(h, t, nil, nil)
	al.Construct(t, nil, h, nil)
	return &List{al, h, t, 0}
}