package main

import (
	"fmt"
	pb "github/ozzy/grpcgo/protoOut"
	"github/ozzy/grpcgo/protofunc/complex"
	"github/ozzy/grpcgo/protofunc/enum"
	"github/ozzy/grpcgo/protofunc/maps"
	"github/ozzy/grpcgo/protofunc/oneofs"
	"github/ozzy/grpcgo/protofunc/simple"
	"github/ozzy/grpcgo/protofunc/tojson"
	"github/ozzy/grpcgo/protofunc/wrtobin"
	"google.golang.org/protobuf/proto"
	"reflect"
)

var path = "protoOutBin/simple.bin"

func main() {
	fmt.Println(simple.DoSimple())
	fmt.Println(complex.DoComplex())
	fmt.Println(enum.DoEnum())
	oneofs.DoOneOfs(&pb.Result_Id{Id: 42})
	oneofs.DoOneOfs(&pb.Result_Message{Message: "A message"})
	fmt.Print(maps.DoMap())
	wrtobin.WriteToFile(path, simple.DoSimple())
	message := &pb.Simple{}
	wrtobin.ReadFromFile(path, message)
	fmt.Println(message)
	s := tojson.ToJson(message)
	message2 := &pb.Simple{}
	tojson.FromJson(s, message2)
	fmt.Println(message2)
	js := doToJson(simple.DoSimple())
	message3 := doFromJson(js, reflect.TypeOf(pb.Simple{}))
	fmt.Println(js)
	fmt.Println(message3)
}

func doToJson(p proto.Message) string {
	jsonString := tojson.ToJson(p)
	return jsonString
}

func doFromJson(jsonString string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	tojson.FromJson(jsonString, message)
	return message
}
