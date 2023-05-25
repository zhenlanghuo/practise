package main

import (
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/varint"
)

func main() {
	//ord.SizeString("")
	//ord.MarshalString()
	//ord.UnmarshalString()
	//ord.UnmarshalMap()
	//
	//unsafe.SizeString()
	//unsafe.MarshalString()
	//unsafe.UnmarshalString()

	var (
		sl = []int{1, 2, 3, 4, 5}
		m  = mus.MarshalerFn[int](varint.MarshalInt) // Implementation of the
		// mus.Marshaler interface for slice elements.
		u = mus.UnmarshalerFn[int](varint.UnmarshalInt) // Implementation of the
		// mus.Unmarshaler interface for slice elements.
		s = mus.SizerFn[int](varint.SizeInt) // Implementation of the mus.Sizer
		// interface for slice elements.
		size = ord.SizeSlice[int](sl, s)
		bs   = make([]byte, size)
	)
	n := ord.MarshalSlice[int](sl, m, bs)
	sl, n, err := ord.UnmarshalSlice[int](u, bs)

	ord.MarshalPtr()

}
