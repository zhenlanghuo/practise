package main

import (
	muscom "github.com/mus-format/mus-common-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/varint"
)

type PtrMarshaler[T any] interface {
	MarshalMUS(t *T, bs []byte) (n int)
}

type PtrMarshalerFn[T any] func(t *T, bs []byte) (n int)

func (fn PtrMarshalerFn[T]) MarshalMUS(t *T, bs []byte) (n int) {
	return fn(t, bs)
}

func MarshalPtrExtern[T any](v *T, m PtrMarshaler[T], bs []byte) (n int) {
	if v == nil {
		bs[0] = muscom.NilFlag
		n = 1
		return
	}
	bs[0] = muscom.NotNilFlag
	return 1 + m.MarshalMUS(v, bs[1:])
}

type PtrUnmarshaler[T any] interface {
	UnmarshalMUS(bs []byte) (t *T, n int, err error)
}

type PtrUnmarshalerFn[T any] func(bs []byte) (t *T, n int, err error)

func (fn PtrUnmarshalerFn[T]) UnmarshalMUS(bs []byte) (t *T, n int, err error) {
	return fn(bs)
}

func UnmarshalPtrExtern[T any](u PtrUnmarshaler[T], bs []byte) (v *T, n int,
	err error) {
	if len(bs) < 1 {
		err = mus.ErrTooSmallByteSlice
		return
	}
	if bs[0] == muscom.NilFlag {
		n = 1
		return
	}
	if bs[0] != muscom.NotNilFlag {
		err = muscom.ErrWrongFormat
		return
	}
	k, n, err := u.UnmarshalMUS(bs[1:])
	if err != nil {
		n = 1 + n
		return
	}
	return k, 1 + n, err
}

type PtrSizer[T any] interface {
	SizeMUS(t *T) (size int)
}

type PtrSizerFn[T any] func(t *T) (size int)

func (fn PtrSizerFn[T]) SizeMUS(t *T) (size int) {
	return fn(t)
}

func SizePtrExtern[T any](v *T, s PtrSizer[T]) (size int) {
	if v != nil {
		return 1 + s.SizeMUS(v)
	}
	return 1
}

func SizePtrSlice[T any](v []*T, s PtrSizer[T]) (size int) {
	size = varint.SizeInt(len(v))
	for i := 0; i < len(v); i++ {
		size += s.SizeMUS(v[i]) + 1
	}
	return
}

func MarshalPtrSlice[T any](v []*T, m PtrMarshaler[T], bs []byte) (n int) {
	n = varint.MarshalInt(len(v), bs)
	for _, e := range v {
		n += MarshalPtrExtern[T](e, m, bs[n:])
	}
	return
}

func UnmarshalPtrSlice[T any](u PtrUnmarshaler[T], bs []byte) (v []*T, n int,
	err error) {
	return UnmarshalValidPtrSlice(nil, u, nil, nil, bs)
}

func UnmarshalValidPtrSlice[T any](maxLength muscom.Validator[int],
	u PtrUnmarshaler[T],
	vl muscom.Validator[T],
	sk mus.Skipper,
	bs []byte,
) (v []*T, n int, err error) {
	length, n, err := varint.UnmarshalInt(bs)
	if err != nil {
		return
	}
	if length < 0 {
		err = muscom.ErrNegativeLength
		return
	}
	var (
		n1   int
		err1 error
		i    int
		e    *T
	)
	if maxLength != nil {
		if err = maxLength.Validate(length); err != nil {
			goto SkipRemainingBytes
		}
	}
	v = make([]*T, length)
	for i = 0; i < length; i++ {
		e, n1, err = UnmarshalPtrExtern[T](u, bs[n:])
		n += n1
		if err != nil {
			return
		}
		if vl != nil && e != nil {
			if err = vl.Validate(*e); err != nil {
				i++
				goto SkipRemainingBytes
			}
		}
		v[i] = e
	}
	return
SkipRemainingBytes:
	if sk == nil {
		return
	}
	n1, err1 = skipRemainingSlice(i, length, sk, bs[n:])
	n += n1
	if err1 != nil {
		err = err1
	}
	return
}

func skipRemainingSlice(from int, length int, sk mus.Skipper, bs []byte) (n int,
	err error) {
	var n1 int
	for i := from; i < length; i++ {
		n1, err = sk.SkipMUS(bs[n:])
		n += n1
		if err != nil {
			return
		}
	}
	return
}
