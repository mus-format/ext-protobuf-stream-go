package exts

import (
	muss "github.com/mus-format/mus-stream-go"
	strops "github.com/mus-format/mus-stream-go/options/string"
	"github.com/mus-format/mus-stream-go/ord"
	"github.com/mus-format/mus-stream-go/unsafe"
	"github.com/mus-format/mus-stream-go/varint"
)

var (
	LenSer       = lenSer{}
	String       = ord.NewStringSer(strops.WithLenSer(LenSer))
	UnsafeString = unsafe.NewStringSer(strops.WithLenSer(LenSer))
)

// NewValidStringProtobuf returns a new valid string serializer.
func NewValidStringProtobuf(ops ...strops.SetOption) muss.Serializer[string] {
	ops = append(ops, strops.WithLenSer(LenSer))
	return ord.NewValidStringSer(ops...)
}

// NewValidStringUnsafeProtobuf returns a new valid string serializer.
func NewValidStringUnsafeProtobuf(ops ...strops.SetOption) muss.Serializer[string] {
	ops = append(ops, strops.WithLenSer(LenSer))
	return unsafe.NewValidStringSer(ops...)
}

// lenSer implements the mus.Serializer interface for length.
type lenSer struct{}

func (lenSer) Marshal(v int, w muss.Writer) (n int, err error) {
	return varint.PositiveInt32.Marshal(int32(v), w)
}

func (lenSer) Unmarshal(r muss.Reader) (v int, n int, err error) {
	v32, n, err := varint.PositiveInt32.Unmarshal(r)
	v = int(v32)
	return
}

func (lenSer) Size(v int) (size int) {
	return varint.PositiveInt32.Size(int32(v))
}

func (lenSer) Skip(r muss.Reader) (n int, err error) {
	return varint.PositiveInt32.Skip(r)
}
