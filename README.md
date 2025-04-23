# ext-protobuf-stream-go
Provides a [mus-stream-go](https://github.com/mus-format/mus-go) serializer 
extension for the Protobuf format.

Includes the `MarshallerProtobuf` interface, which represents a type that can 
marshal itself into the Protobuf format. Also includes the `MarshallerTypedProtobuf` 
interface, intended for use with [DTS](https://github.com/mus-format/mus-stream-dts-go).

Contains serializers for string, slice and timestamp types.