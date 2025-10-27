package ext

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-stream-go"
	slops "github.com/mus-format/mus-stream-go/options/slice"
	"github.com/mus-format/mus-stream-go/varint"
)

// NewSliceSer returns a new slice serializer with the given element serializer.
func NewSliceSer[T any](elemSer mus.Serializer[T]) sliceSer[T] {
	return sliceSer[T]{elemSer}
}

// NewValidSliceSer returns a new valid slice serializer.
func NewValidSliceSer[T any](elemSer mus.Serializer[T],
	ops ...slops.SetOption[T],
) validSliceProtobuf[T] {
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
		sliceSer: NewSliceSer(elemSer),
		lenVl:    lenVl,
		elemVl:   elemVl,
	}
}

// sliceSer implements the mus.Serializer interface for slices.
type sliceSer[T any] struct {
	elemProtobuf mus.Serializer[T]
}

func (s sliceSer[T]) Marshal(sl []T, w mus.Writer) (n int, err error) {
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
		for i := range len(sl) {
			n1, err = s.elemProtobuf.Marshal(sl[i], w)
			n += n1
			if err != nil {
				return
			}
		}
	}
	return
}

func (s sliceSer[T]) Unmarshal(r mus.Reader) (sl []T, n int, err error) {
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

func (s sliceSer[T]) Size(sl []T) (size int) {
	size = s.size(sl)
	return size + varint.PositiveInt.Size(size)
}

func (s sliceSer[T]) Skip(r mus.Reader) (n int, err error) {
	l, n, err := varint.PositiveInt.Unmarshal(r)
	if err != nil {
		return
	}
	n += l
	return
}

func (s sliceSer[T]) size(sl []T) (size int) {
	for i := range sl {
		size += s.elemProtobuf.Size(sl[i])
	}
	return
}

// -----------------------------------------------------------------------------

type validSliceProtobuf[T any] struct {
	sliceSer[T]
	lenVl  com.Validator[int]
	elemVl com.Validator[T]
}

func (s validSliceProtobuf[T]) Unmarshal(r mus.Reader) (sl []T, n int,
	err error,
) {
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
