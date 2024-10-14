package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ProtobufToJSON converts protocol buffer message to JSON string
func ProtobufToJSON(message proto.Message) (string, error) {
	marshaler := protojson.MarshalOptions{
		UseEnumNumbers:    false,
		EmitDefaultValues: true,
		Indent:            "  ",
		UseProtoNames:     true,
	}

	ret, err := marshaler.Marshal(message)
	return string(ret), err
}

// JSONToProtobufMessage converts JSON string to protocol buffer message
func JSONToProtobufMessage(data string, message proto.Message) error {
	return protojson.UnmarshalOptions{}.Unmarshal([]byte(data), message)
}
