package exts

import (
	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	slops "github.com/mus-format/mus-stream-go/options/slice"
	"github.com/mus-format/mus-stream-go/varint"
)

// NewSliceSer returns a new slice serializer with the given element serializer.
func NewSliceProtobuf[T any](elemProtobuf muss.Serializer[T]) sliceProtobuf[T] {
	return sliceProtobuf[T]{elemProtobuf}
}

// NewValidSliceSer returns a new valid slice serializer.
func NewValidSliceProtobuf[T any](elemProtobuf muss.Serializer[T],
	ops ...slops.SetOption[T]) validSliceProtobuf[T] {
	o := slops.Options[T]{}
	slops.Apply(ops, &o)

	var (
		lenVl  com.Validator[int]
		elemVl com.Validator[T]
	)
	if o.LenVl != nil {
		lenVl = o.LenVl
	}
	if o.ElemVl != nil {
		elemVl = o.ElemVl
	}
	return validSliceProtobuf[T]{
		sliceProtobuf: NewSliceProtobuf(elemProtobuf),
		lenVl:         lenVl,
		elemVl:        elemVl,
	}
}

// sliceProtobuf implements the mus.Serializer interface for slices.
type sliceProtobuf[T any] struct {
	elemProtobuf muss.Serializer[T]
}

func (s sliceProtobuf[T]) Marshal(sl []T, w muss.Writer) (n int, err error) {
	var (
		n1     int
		length = len(sl)
	)
	if length > 0 {
		n1, err = varint.PositiveInt.Marshal(s.size(sl), w)
		n += n1
		if err != nil {
			return
		}
		for i := 0; i < len(sl); i++ {
			n1, err = s.elemProtobuf.Marshal(sl[i], w)
			n += n1
			if err != nil {
				return
			}
		}
	}
	return
}

func (s sliceProtobuf[T]) Unmarshal(r muss.Reader) (sl []T, n int, err error) {
	var (
		n1 int
		e  T
	)
	sl = []T{}
	size, n, err := varint.PositiveInt.Unmarshal(r)
	if err != nil {
		return
	}
	for n < size {
		e, n1, err = s.elemProtobuf.Unmarshal(r)
		n += n1
		if err != nil {
			return
		}
		sl = append(sl, e)
	}
	return
}

func (s sliceProtobuf[T]) Size(sl []T) (size int) {
	size = s.size(sl)
	return size + varint.PositiveInt.Size(size)
}

func (s sliceProtobuf[T]) Skip(r muss.Reader) (n int, err error) {
	l, n, err := varint.PositiveInt.Unmarshal(r)
	if err != nil {
		return
	}
	n += l
	return
}

func (s sliceProtobuf[T]) size(sl []T) (size int) {
	for i := range sl {
		size += s.elemProtobuf.Size(sl[i])
	}
	return
}

// -----------------------------------------------------------------------------

type validSliceProtobuf[T any] struct {
	sliceProtobuf[T]
	lenVl  com.Validator[int]
	elemVl com.Validator[T]
}

func (s validSliceProtobuf[T]) Unmarshal(r muss.Reader) (sl []T, n int, err error) {
	var (
		n1 int
		e  T
	)
	sl = []T{}
	size, n, err := varint.PositiveInt.Unmarshal(r)
	if err != nil {
		return
	}
	if s.lenVl != nil {
		if err = s.lenVl.Validate(size); err != nil {
			return
		}
	}
	for n < size {
		e, n1, err = s.elemProtobuf.Unmarshal(r)
		n += n1
		if err != nil {
			return
		}
		if s.elemVl != nil {
			if err = s.elemVl.Validate(e); err != nil {
				return
			}
		}
		sl = append(sl, e)
	}
	return
}
