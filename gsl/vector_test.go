package gsl_test

import (
	gsl "github.com/hackbutteers/wgtl/gsl"
	"testing"
	"fmt"
)

func TestVector(T *testing.T) {
	fmt.Println("vector test start")
	v := gsl.NewVector()
	v.PushBack(1)
	v.PushBack(2)
	v.PushBack(3)
	v.PushBack(2)


	rit := v.FindLastOf(2, &gsl.IntComparator{})

	fmt.Println(rit.Index())
	fmt.Println("")
	v.Sort(&gsl.IntComparator{})
	for it := v.Begin(); !it.Equal(v.End()); it.Next() {
		fmt.Println(it.Value())
	}
	fmt.Println("")

	fmt.Println("vector test end")

}