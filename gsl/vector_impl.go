package gsl


func (v *Vector)grow(needsize int) {
	s := v.end - v.begin
	if s + needsize < v.cap {
		return
	}
	var ncap int
	if v.cap > 1024 * 1024 {
		ncap = v.cap + int(float32(needsize) * growFactor)
	} else {
		ncap = v.cap * 2 + needsize
	}
	nb := v.alloc.Allocate(ncap).([]interface{})
	cb := v.data[v.begin:v.end]
	copy(nb, cb)
	v.alloc.Deallocate(interface{}(v.data), v.cap)
	v.data = nb
	v.cap = ncap
}
func (v *Vector)PushBack(values...interface{}) {
	size := len(values)
	v.grow(size)
	for i := 0; i < size; i++ {
		v.data[v.end] = values[i]
		v.end++
	}
}
func (v *Vector)PopBack() interface{} {
	if v.end <= v.begin {
		return nil
	}
	v.end--
	return v.data[v.end]
}
func (v *Vector)PopFront() interface{} {
	if v.begin >= v.end {
		return nil
	}
	t := v.begin
	v.begin++
	return v.data[t]
}

func(v *Vector)withinRange(index int) bool {
	if index <0 || index > v.end - v.begin {
		return false
	}
	return true
}

func (v *Vector)Get(index int) interface{} {
	if !v.withinRange(index) {
		return nil
	}
	return v.data[v.begin + index]
}
func (v *Vector)Remove(index int) {
	if !v.withinRange(index) {
		return
	}
	v.data[index] = nil
	copy(v.data[index:], v.data[index+1:v.end])
	v.end--
}

func (v *Vector)RemoveItr(it *VectorIterator) {
	i := it.Index()
	v.Remove(i)
}

func (v *Vector)RemoveRitr(rit *VectorReverseIterator) {
	ri := rit.Index()
	i := v.end - v.begin
	i = i - ri
	v.Remove(i)
}

func (v *Vector)FindFirstOf(value interface{}, comp Comparator) *VectorIterator {
	for i := v.begin; i < v.end; i++ {
		if comp.Compare(v.data[i], value) == 0 {
			return &VectorIterator{ i - v.begin, v}
		}
	}
	return v.End()
}
func (v *Vector)FindLastOf(value interface{}, comp Comparator) *VectorReverseIterator {
	for i := v.end - 1; i > v.begin - 1; i-- {
		if comp.Compare(v.data[i], value) == 0 {
			//fmt.Println(i)
			return &VectorReverseIterator{i, v}
		}
	}
	return v.Rend()
}

func (v *Vector)Reserve(size int) {
	v.grow(size)
}

func (v *Vector)Sort(comp Comparator) {
	b := v.data[v.begin:v.end]
	Sort(b,  comp)
}

func (v *Vector)Swap(i,j int) {
	v.data[i], v.data[j] = v.data[j], v.data[i]
}

func (v *Vector)Insert(index int, values...interface{}) {
	vs := len(values)
	v.grow(vs)
	v.end += vs

	for i := 0; i  < vs; i++ {
		t := v.end - 1 -i
		s := t - vs
		v.data[t] = v.data[s]
	}
	for i, value := range values {
		v.data[index+i] = value
	}
}


func (v *Vector)Clear() {
	v.begin = prefix
	v.end = prefix
}

func (v *Vector)Size() int{
	return v.end - v.begin
}

func (v *Vector)Empty() bool{
	return v.end - v.begin == 0
}

func (v *Vector)Begin() *VectorIterator {
	return &VectorIterator{v.begin, v}
}

func (v *Vector)End() *VectorIterator {
	return &VectorIterator{v.end, v}
}
func (v *Vector)Rbegin() *VectorReverseIterator {
	return &VectorReverseIterator{0, v}	
}
func (v *Vector) Rend() *VectorReverseIterator {
	return &VectorReverseIterator{v.begin - 1, v}	
}