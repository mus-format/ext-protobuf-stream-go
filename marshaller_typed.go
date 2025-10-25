package ext

import "github.com/mus-format/mus-stream-go"

// MarshallerProtobuf interface wraps the MarshalProtobuf and SizeProtobuf
// methods. It is intended for use with DTS.
type MarshallerTypedProtobuf interface {
	MarshalTypedProtobuf(w mus.Writer) (n int, err error)
	SizeTypedProtobuf() (size int)
}
