package gsl

type Allocator interface {
	Allocate(size int) interface{}
	Deallocate(v interface{}, size int)
	Construct(v interface{}, args...interface{})
}