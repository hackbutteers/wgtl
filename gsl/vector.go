package gsl



type VectorInterface interface {
	PushBack(v...interface{})
	PopBack() interface{}
	PopFront() interface{}
	Get(index int) interface{}
	Remove(index int)
	RemoveItr(it *VectorIterator)
	RemoveRitr(rit *VectorReverseIterator)
	FindFirstOf(v interface{}, comp Comparator) *VectorIterator
	FindLastOf(v interface{}, comp Comparator) *VectorReverseIterator
	Reserve(size int)
	Sort(comp Comparator)
	Swap(i,j int)
	Insert(i int, v...interface{})
	ContainerInterface

	Begin() *VectorIterator
	End() *VectorIterator
	Rbegin() *VectorReverseIterator
	Rend() *VectorReverseIterator

}

type Vector struct {
	data []interface{}
	alloc Allocator
	begin int
	end   int
	cap  int
}

const (
	initSize int= 16
	prefix   int= 8
	growFactor float32 = 2
)

type VectorIterator struct {
	index int
	owner *Vector
}

type VectorReverseIterator struct {
	index int
	owner *Vector
}

type DefaultVectorAllocator struct {}

func (a *DefaultVectorAllocator)Allocate(s int) interface{} {
	r := make([]interface{}, s, s)
	return r
}
func (a *DefaultVectorAllocator)Deallocate(v interface{}, size int) {

}
func (a *DefaultVectorAllocator)Construct(v interface{}, args...interface{}) {

}

func NewVector() *Vector{
	return NewAllocateVector(nil)
}


func NewAllocateVector(a Allocator) *Vector {
	al := &DefaultVectorAllocator{}
	if a != nil {
		al = a.(*DefaultVectorAllocator)
	} 
	b := al.Allocate(initSize)
	return &Vector{b.([]interface{}), al, 0, 0 , 8}
}


func (i *VectorIterator)Next() bool {
	if i.index + i.owner.begin >= i.owner.end {
		return false
	}
	i.index++
	return true
}

func (i *VectorIterator)Prev() bool {
	if i.index <= 0 {
		return false
	}
	i.index--
	return true
}

func (i *VectorIterator)Value() interface{} {
	if i.index < 0 || i.index + i.owner.begin >= i.owner.end {
		return nil
	}
	return i.owner.data[i.index + i.owner.begin]
}

func (i *VectorIterator)Index() int {
	return i.index
}

func (i *VectorIterator)First() bool {
	i.index = 0
	return true
}

func (i *VectorIterator)Assign(v interface{}) {
	i.owner.data[i.index + i.owner.begin] = v
}

func (i *VectorIterator)Equal(t BidirectIterator) bool {
	return i.index == t.Index()
}

func (i *VectorReverseIterator)Next() bool {
	if i.owner.end - i.index - 1 < i.owner.begin {
		return false
	}
	i.index--
	return true
}

func (i *VectorReverseIterator)Prev() bool {
	if i.index >= i.owner.end  {
		return false
	}
	i.index++
	return true
}

func (i *VectorReverseIterator)Value() interface{} {
	if i.index < 0 || i.owner.end - i.index < i.owner.begin {
		return nil
	}
	return i.owner.data[i.owner.end - i.index -1]
}

func (i *VectorReverseIterator)Index() int {
	return i.index
}

func (i *VectorReverseIterator)First() bool {
	i.index = i.owner.end - 1 - i.owner.begin
	return true
}

func (i *VectorReverseIterator)Assign(v interface{}) {
	i.owner.data[i.index + i.owner.begin] = v
}

func (i *VectorReverseIterator)Equal(t BidirectIterator) bool {
	return i.index == t.Index()
}