package mock

import (
	muss "github.com/mus-format/mus-stream-go"
	"github.com/ymz-ncnk/mok"
)

type MarshalProtobufFn func(w muss.Writer) (n int, err error)
type SizeProtobufFn func() (size int)

func NewMarshallerProtobuf() MarshallerProtobuf {
	return MarshallerProtobuf{mok.New("MarshallerProtobuf")}
}

type MarshallerProtobuf struct {
	*mok.Mock
}

func (m MarshallerProtobuf) RegisterMarshalProtobuf(fn MarshalProtobufFn) MarshallerProtobuf {
	m.Register("MarshalProtobuf", fn)
	return m
}

func (m MarshallerProtobuf) RegisterSizeProtobuf(fn SizeProtobufFn) MarshallerProtobuf {
	m.Register("SizeeProtobuf", fn)
	return m
}

func (m MarshallerProtobuf) MarshalProtobuf(w muss.Writer) (n int, err error) {
	result, err := m.Call("MarshalProtobuf", w)
	if err != nil {
		panic(err)
	}
	n = result[0].(int)
	err, _ = result[1].(error)
	return
}

func (m MarshallerProtobuf) SizeProtobuf() (size int) {
	result, err := m.Call("SizeProtobuf")
	if err != nil {
		panic(err)
	}
	return result[0].(int)
}
