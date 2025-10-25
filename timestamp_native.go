package ext

import (
	"github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/varint"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var TimestampNativeProtobuf = timestampNativeProtobuf{}

var (
	secondsFieldTag = protowire.EncodeTag(1, protowire.VarintType)
	nanosFieldTag   = protowire.EncodeTag(2, protowire.VarintType)
)

// timestampNativeProtobuf implements the mus.Serializer interface for
// timestamppb.Timestamp.
type timestampNativeProtobuf struct{}

func (s timestampNativeProtobuf) Marshal(tm *timestamppb.Timestamp,
	w mus.Writer,
) (n int, err error) {
	var (
		n1   int
		size = s.size(tm)
	)
	if size > 0 {
		n1, err = varint.PositiveInt.Marshal(size, w)
		n += n1
		if err != nil {
			return
		}
		if tm.Seconds != 0 {
			n1, err = varint.Uint64.Marshal(secondsFieldTag, w)
			n += n1
			if err != nil {
				return
			}
			n1, err = varint.PositiveInt64.Marshal(tm.Seconds, w)
			n += n1
			if err != nil {
				return
			}
		}
		if tm.Nanos != 0 {
			n1, err = varint.Uint64.Marshal(nanosFieldTag, w)
			n += n1
			if err != nil {
				return
			}
			n1, err = varint.PositiveInt32.Marshal(tm.Nanos, w)
			n += n1
			if err != nil {
				return
			}
		}
	}
	return
}

func (timestampNativeProtobuf) Unmarshal(r mus.Reader) (
	tm *timestamppb.Timestamp, n int, err error,
) {
	size, _, err := varint.PositiveInt.Unmarshal(r)
	if err != nil {
		return
	}
	var (
		n1  int
		tag uint64
	)
	tm = &timestamppb.Timestamp{}
	for n < size {
		tag, n1, err = varint.Uint64.Unmarshal(r)
		n += n1
		if err != nil {
			return
		}
		switch tag {
		case secondsFieldTag:
			tm.Seconds, n1, err = varint.PositiveInt64.Unmarshal(r)
		case nanosFieldTag:
			tm.Nanos, n1, err = varint.PositiveInt32.Unmarshal(r)
		}
		n += n1
		if err != nil {
			return
		}
	}
	return
}

func (s timestampNativeProtobuf) Size(tm *timestamppb.Timestamp) (size int) {
	size = s.size(tm)
	return size + varint.PositiveInt.Size(size)
}

func (s timestampNativeProtobuf) Skip(r mus.Reader) (n int, err error) {
	size, _, err := varint.PositiveInt.Unmarshal(r)
	if err != nil {
		return
	}
	var (
		n1  int
		tag uint64
	)
	for n < size {
		tag, n1, err = varint.Uint64.Unmarshal(r)
		n += n1
		if err != nil {
			return
		}
		switch tag {
		case secondsFieldTag:
			n1, err = varint.PositiveInt64.Skip(r)
		case nanosFieldTag:
			n1, err = varint.PositiveInt32.Skip(r)
		}
		n += n1
		if err != nil {
			return
		}
	}
	return
}

func (s timestampNativeProtobuf) size(tm *timestamppb.Timestamp) (size int) {
	if tm.Seconds != 0 {
		size += varint.Uint64.Size(secondsFieldTag)
		size += varint.PositiveInt64.Size(tm.Seconds)
	}
	if tm.Nanos != 0 {
		size += varint.Uint64.Size(nanosFieldTag)
		size += varint.PositiveInt32.Size(tm.Nanos)
	}
	return
}
