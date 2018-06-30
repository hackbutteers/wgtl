package gsl_test

import (
	gsl "github.com/hackbutteers/wgtl/gsl"
	"testing"
	"fmt"
)

func TestList(t *testing.T) {
	l := gsl.NewList()
	l.PushFront(1)
	l.PushFront(2)
	l.PushFront(3)
	l.PushBack(10)
	l.PushBack(11)

	fmt.Println(l.Size())
	
	for it := l.Begin(); !it.Equal(l.End()); it.Next() {
		fmt.Println("list elemnt index:", it.Index(), "element value:", it.Value())
	}
	fmt.Println("")
	for it := l.Rbegin(); !it.Equal(l.Rend()); it.Next() {
		fmt.Println("list elemnt index:", it.Index(), "element value:", it.Value())
	}
	comp := &gsl.IntComparator{}
	it := l.FindValue(3, comp)
	l.RemoveItr(it)
	fmt.Println("")
	for it := l.Begin(); !it.Equal(l.End()); it.Next() {
		fmt.Println("list elemnt index:", it.Index(), "element value:", it.Value())
	}
	l.Sort(comp)
	fmt.Println("")
	for it := l.Begin(); !it.Equal(l.End()); it.Next() {
		fmt.Println("list elemnt index:", it.Index(), "element value:", it.Value())
	}
	
}