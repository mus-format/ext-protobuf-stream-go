# ext-protobuf-stream-go

Provides a [mus-stream-go](https://github.com/mus-format/mus-stream-go) serializer
extension for the Protobuf format.

This package includes:

- `MarshallerProtobuf` — an interface for types that can marshal themselves
  into the Protobuf format.
- `MarshallerTypedProtobuf` — an interface for types that support typed Protobuf
  serialization (designed for use with [DTS](https://github.com/mus-format/dts-go)).
- Serializers for string, slice and timestamp types.
