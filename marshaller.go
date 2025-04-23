package exts

import muss "github.com/mus-format/mus-stream-go"

// MarshallerProtobuf interface wraps MarhsalProtobuf and SizeProtobuf methods.
type MarshallerProtobuf interface {
	MarshalProtobuf(w muss.Writer) (n int, err error)
	SizeProtobuf() (size int)
}
