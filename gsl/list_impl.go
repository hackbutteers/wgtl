package gsl
// allocator
func (al *ListDefaultAllocator)Allocate() interface{} {
	return &Lnode{}
}

func (al *ListDefaultAllocator)Deallocate(v interface{}) {
	
}

func (al *ListDefaultAllocator)Construct(v interface{}, args...interface{}) {
	va := v.(*Lnode)
	va.next = args[0].(*Lnode)
	va.prev = args[1].(*Lnode)
	if len(args) == 3 {
		va.value = args[2]
	}
}

//member functions

func (l *List) PushFront(v interface{}) {
	ln := l.alloc.Allocate().(*Lnode)
	l.alloc.Construct(ln, l.head.next, l.head, v)
	defer l.alloc.Deallocate(ln)
	l.head.next.prev = ln
	l.head.next = ln
	l.size = l.size + 1
}

func (l *List) PushBack(v interface{}) {
	ln := l.alloc.Allocate().(*Lnode)
	l.alloc.Construct(ln, l.tail, l.tail.prev, v)
	defer l.alloc.Deallocate(ln)
	l.tail.prev.next = ln
	l.tail.prev = ln
	l.size = l.size + 1
}

func (l *List) Front() interface{} {
	if l.size == 0 {
		return nil
	}
	return l.head.next
}

func (l *List) PopFront() interface{} {
	if l.size == 0 {
		return nil
	}	
	r := l.head.next
	r.next.prev = l.head
	l.head.next = r.next
	l.size = l.size - 1
	return r
}

func (l *List) PopBack() interface{} {
	if l.size == 0 {
		return nil
	}	
	r := l.tail.prev
	r.prev.next = l.tail
	l.tail.prev = r.prev
	l.size = l.size - 1
	return r
}

func (l *List) back() interface{} {
	if l.size == 0 {
		return nil
	}
	return l.tail.prev
}

func (l *List)getInternal(index int) *Lnode {
	if !l.indexValid(index) {
		return nil
	}
	var ln *Lnode
	if index < l.size/2 {
		ln = l.head.next
		for i := 0; i < index; i++ {
			ln = ln.next
		}
	} else {
		ln = l.tail.prev
		for i := l.size - 1; i > index; i-- {
			ln = ln.prev
		}
	}
	return ln
}

func swapNode(n1 *Lnode, n2 *Lnode) {
	v := n1.value
	n1.value = n2.value
	n2.value = v
}
func (l *List)Swap(index1 int, index2 int) bool {
	ln1 := l.getInternal(index1)
	ln2 := l.getInternal(index2)
	if ln1 == nil || ln2 == nil {
		return false
	}
	swapNode(ln1, ln2)
	return true;
}

func (l *List)RemoveIndex(index int) {
	ln := l.getInternal(index) 
	if ln == nil {
		return
	}
	ln.prev.next = ln.next
	ln.next.prev = ln.prev
	l.size--
}

func (l *List)RemoveItr(it *ListIterator) *ListIterator {
	rItr := &ListIterator{it.Index() + 1, it.node.next, l}
	n := it.node
	n.prev.next = n.next
	n.next.prev = n.prev
	l.size--
	return rItr
}

func  (l *List)RemoveValue(v interface{}, comp Comparator) {
	it := l.FindValue(v, comp)
	_ = l.RemoveItr(it)
}


func (l *List)indexValid(index int) bool {
	return index >= 0 && index < l.size
}
func (l *List)Insert(index int, values ...interface{}) {
	if !l.indexValid(index) {
		return 
	}
	l.size = l.size + len(values)
	ln := l.getInternal(index)
	lnNext := ln.next
	var in *Lnode
	for _, e := range values {
		in = l.alloc.Allocate().(*Lnode)
		ln.next = in
		in.prev = ln
		in.value = e
		ln = ln.next
	}
	ln.next = lnNext
	ln.prev = ln
}

func (l *List)Begin() *ListIterator {
	return &ListIterator{0, l.head.next, l}
}

func (l *List)End() *ListIterator {
	return &ListIterator{l.size, l.tail, l}
}

func (l *List)Rbegin() *ListReverseIterator {
	return &ListReverseIterator{l.size - 1, l.tail.prev, l}
}

func (l *List)Rend() *ListReverseIterator {
	return &ListReverseIterator{-1, l.head, l}
}


func (l *List)Slice() []interface{} {
	slice := make([]interface{}, 0, l.size)

	for ln := l.head.next; ln != l.tail; ln = ln.next {
		slice = append(slice, ln.value)
	}
	return slice
}

func (l *List)Sort(comp Comparator) {
	if l.size < 2 {
		return
	}

	values := l.Slice()
	Sort(values, comp)

	l.Clear()

	l.Add(values...)
}

func (l *List)Clear() {
	l.head.next = l.tail
	l.tail.prev = l.head
	l.size = 0
}

func (l * List) Add(v...interface{}) {
	ln := l.tail.prev
	for _, e := range v {
		in := l.alloc.Allocate().(*Lnode)
		ln.next = in
		in.prev = ln
		in.value = e
		ln = ln.next
	}
	ln.next = l.tail
	l.tail.prev = ln
	l.size += len(v)
}

func (l *List)Empty() bool {
	return l.size == 0
}

func (l *List)Size() int {
	return l.size 
}

func (l *List)PushFrontv(v... interface{}) {
	for i := len(v)-1; i >= 0; i-- {
		l.PushFront(v[i])
	}	
}

func (l *List)PushBackv(v...interface{}) {
	l.Add(v...)
}

func (i *ListIterator)Next() bool {
	if i.node == i.list.tail {
		return false
	}
	i.node = i.node.next
	i.index++
	return true
}

func (i *ListIterator)Prev() bool {
	if i.node == i.list.head {
		return false
	}
	i.node = i.node.prev
	i.index--
	return true
}

func (i *ListIterator)Value() interface{} {
	if i.node == i.list.head || i.node == i.list.tail {
		return nil
	}
	return i.node.value
}

func (i *ListIterator)Index() int {
	if i.node == i.list.head || i.node == i.list.tail {
		return -1
	}
	return i.index
}

func (i *ListIterator)Equal(rhs *ListIterator) bool {
	return rhs.list == i.list && rhs.index == i.index
}


func (i *ListReverseIterator)Next() bool {
	if i.node == i.list.head {
		return false
	}
	i.node = i.node.prev
	i.index--
	return true
}

func (i *ListReverseIterator)Prev() bool {
	if i.node == i.list.tail {
		return false
	}
	i.node = i.node.next
	i.index++
	return true
}

func (i *ListReverseIterator)Value() interface{} {
	if i.node == i.list.head || i.node == i.list.tail {
		return nil
	}
	return i.node.value
}

func (i *ListReverseIterator)Index() int {
	if i.node == i.list.head || i.node == i.list.tail {
		return -1
	}
	return i.index
}

func (i *ListReverseIterator)Equal(rhs *ListReverseIterator) bool {
	return rhs.list == i.list && rhs.index == i.index
}

func (l *List)FindIndex(index int) *ListIterator {
	n := l.getInternal(index)
	if n == nil {
		return nil
	}
	return &ListIterator{index, n, l}
}

func (l *List)FindValue(v interface{}, comp Comparator) *ListIterator {
	ln := l.head.next
	index := 0
	for ;ln != l.tail; ln = ln.next {
		r := comp.Compare(ln.value, v)
		if r == 0 {
			return &ListIterator{index, ln, l}
		}
		index++
	}
	return nil
}