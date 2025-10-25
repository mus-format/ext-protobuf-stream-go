package ext

import (
	"github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/varint"
)

type Timestamp struct {
	Seconds int64
	Nanos   int32
}

var TimestampProtobuf = timestampProtobuf{}

type timestampProtobuf struct{}

func (s timestampProtobuf) Marshal(tm Timestamp, w mus.Writer) (n int, err error) {
	var (
		n1   int
		size = s.size(tm)
	)
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
	return
}

func (s timestampProtobuf) Unmarshal(r mus.Reader) (tm Timestamp, n int,
	err error,
) {
	size, n, err := varint.PositiveInt.Unmarshal(r)
	if err != nil {
		return
	}
	var (
		tag uint64
		n1  int
	)
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

func (s timestampProtobuf) Size(tm Timestamp) (size int) {
	size = s.size(tm)
	return size + varint.PositiveInt.Size(size)
}

func (s timestampProtobuf) Skip(r mus.Reader) (n int, err error) {
	size, n, err := varint.PositiveInt.Unmarshal(r)
	if err != nil {
		return
	}
	var (
		tag uint64
		n1  int
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

func (s timestampProtobuf) size(tm Timestamp) (size int) {
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
