package gsl

type Allocator interface {
	Allocate() interface{}
	Deallocate(v interface{})
	Construct(v interface{}, args...interface{})
}