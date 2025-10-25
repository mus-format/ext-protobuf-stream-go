package ext

import "github.com/mus-format/mus-stream-go"

// MarshallerProtobuf interface wraps MarhsalProtobuf and SizeProtobuf methods.
type MarshallerProtobuf interface {
	MarshalProtobuf(w mus.Writer) (n int, err error)
	SizeProtobuf() (size int)
}
