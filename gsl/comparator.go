package gsl

import (
	"time"
)
type Comparator interface{
	Compare(a, b interface{}) int
}

type StringComparator struct {}

func (sc *StringComparator)Compare(a, b interface{}) int {
	s1 := a.(string)
	s2 := b.(string)
	min := len(s2)
	if len(s1) < len(s2) {
		min = len(s1)
	}
	diff := 0
	for i := 0; i < min && diff == 0; i++ {
		diff = int(s1[i]) - int(s2[i])
	}
	if diff == 0 {
		diff = len(s1) - len(s2)
	}
	if diff < 0 {
		return -1
	}
	if diff > 0 {
		return 1
	}
	return 0
}

func intCompare(a, b int64) int{
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

func uintCompare(a, b uint64) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}	
}

type IntegerCompare struct {}

func (ic *IntegerCompare)Compare(a, b interface{}) int {

	switch a.(type) {
	case int8:
		return intCompare(int64(a.(int8)), int64(b.(int8)))
	case int16:
		return intCompare(int64(a.(int16)), int64(b.(int16)))
	case int32:
		return intCompare(int64(a.(int32)), int64(b.(int32)))
	case int64:
		return intCompare(a.(int64), b.(int64))
	case uint8:
		return uintCompare(uint64(a.(uint8)), uint64(b.(uint8)))
	case uint16:
		return uintCompare(uint64(a.(uint16)), uint64(b.(uint16)))
	case uint32:
		return uintCompare(uint64(a.(uint32)), uint64(b.(uint32)))
	case uint64:
		return uintCompare(a.(uint64), b.(uint64))
	default:
		return 0
	}
}

func float32Compare(a, b float32) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}	
}

func float64Compare(a, b float64) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}	
}

type FloatCompare struct {}

func (fc *FloatCompare) Compare(a, b interface{}) int {
	switch a.(type) {
	case float32:
		return float32Compare(a.(float32), b.(float32))
	case float64:
		return float64Compare(a.(float64), b.(float64))
	default:
		return 0
	}
}

type Int8Comparator struct {}

func(ic * Int8Comparator)Compare(a, b interface{}) int {
	return intCompare(int64(a.(int8)), int64(b.(int8)))
}

type Int16Comparator struct {}

func(ic * Int16Comparator)Compare(a, b interface{}) int {
	return intCompare(int64(a.(int16)), int64(b.(int16)))
}

type Int32Comparator struct {}

func(ic * Int32Comparator)Compare(a, b interface{}) int {
	return intCompare(int64(a.(int32)), int64(b.(int32)))
}

type IntComparator struct {}

func(ic * IntComparator)Compare(a, b interface{}) int {
	return intCompare(int64(a.(int)), int64(b.(int)))
}


type Int64Comparator struct {}

func(ic * Int64Comparator)Compare(a, b interface{}) int {
	return intCompare(int64(a.(int64)), int64(b.(int64)))
}

type Uint8Comparator struct {}

func(ic * Uint8Comparator)Compare(a, b interface{}) int {
	return uintCompare(uint64(a.(uint8)), uint64(b.(uint8)))
}

type Uint16Comparator struct {}

func(ic * Uint16Comparator)Compare(a, b interface{}) int {
	return uintCompare(uint64(a.(uint16)), uint64(b.(uint16)))
}

type Uint32Comparator struct {}

func(ic * Uint32Comparator)Compare(a, b interface{}) int {
	return uintCompare(uint64(a.(uint32)), uint64(b.(uint32)))
}

type Uint64Comparator struct {}

func(ic * Uint64Comparator)Compare(a, b interface{}) int {
	return uintCompare(uint64(a.(uint64)), uint64(b.(uint64)))
}

type Float32Comparator struct {}

func(ic * Float32Comparator)Compare(a, b interface{}) int {
	return float32Compare(a.(float32), b.(float32))
}

type Float64Comparator struct {}

func(ic * Float64Comparator)Compare(a, b interface{}) int {
	return float64Compare(a.(float64), b.(float64))
}

type RuneComparator struct {}

func (ic * RuneComparator)Compare(a, b interface{}) int {
	aAsserted := a.(rune)
	bAsserted := b.(rune)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

type TimeComparator struct {}

func (tc *TimeComparator)Compare(a, b interface{}) int {
	aAsserted := a.(time.Time)
	bAsserted := b.(time.Time)

	switch {
	case aAsserted.After(bAsserted):
		return 1
	case aAsserted.Before(bAsserted):
		return -1
	default:
		return 0
	}
}