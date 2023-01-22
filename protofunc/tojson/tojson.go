package tojson

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"log"
)

func ToJson(pb proto.Message) string {
	option := protojson.MarshalOptions{
		Multiline: true,
	}
	out, err := option.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}
	return string(out)
}

func FromJson(in string, pb proto.Message) {

	if err := protojson.Unmarshal([]byte(in), pb); err != nil {
		log.Fatalln("Couldn't unmarshal from JSON", err)
	}

}
