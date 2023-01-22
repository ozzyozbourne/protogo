package main

import (
	"fmt"
	pb "github/ozzy/grpcgo/protoOut"
	"github/ozzy/grpcgo/protofunc/complex"
	"github/ozzy/grpcgo/protofunc/enum"
	"github/ozzy/grpcgo/protofunc/maps"
	"github/ozzy/grpcgo/protofunc/oneofs"
	"github/ozzy/grpcgo/protofunc/simple"
	"github/ozzy/grpcgo/protofunc/wrtobin"
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
}
