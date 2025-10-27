package ext

import "github.com/mus-format/mus-stream-go"

// MarshallerProtobuf interface wraps MarhsalProtobuf and SizeProtobuf methods.
type MarshallerProtobuf interface {
	MarshalProtobuf(w mus.Writer) (n int, err error)
	SizeProtobuf() (size int)
}

// MarshallerTypedProtobuf interface wraps the MarshalProtobuf and SizeProtobuf
// methods. It is intended for use with DTS.
type MarshallerTypedProtobuf interface {
	MarshalTypedProtobuf(w mus.Writer) (n int, err error)
	SizeTypedProtobuf() (size int)
}
