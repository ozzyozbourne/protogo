package main

import (
	"fmt"
	ps "github/ozzy/grpcgo/protoOut"
	"github/ozzy/grpcgo/protofunc/complex"
)

func doSimple() *ps.Simple {
	return &ps.Simple{
		Id:          42,
		IsSimple:    true,
		Name:        "A name",
		SampleLists: []int32{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
}

func main() {
	//fmt.Println(doSimple())
	fmt.Println(complex.DoComplex())
}
