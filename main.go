package main

import (
	"fmt"
	"github/ozzy/grpcgo/protofunc/complex"
	"github/ozzy/grpcgo/protofunc/simple"
)

func main() {
	fmt.Println(simple.DoSimple())
	fmt.Println(complex.DoComplex())
}
