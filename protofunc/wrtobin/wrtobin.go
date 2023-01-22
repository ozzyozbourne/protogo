package wrtobin

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
)

func WriteToFile(fName string, pb proto.Message) {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes")
	}

	if err = os.WriteFile(fName, out, 0644); err != nil {
		log.Fatalln("Can't write to file")
		return
	}
	fmt.Println("data has been written")
}

func ReadFromFile(fName string, pb proto.Message) {
	in, err := os.ReadFile(fName)
	if err != nil {
		log.Fatalln("Cant read the file", err)
	}

	if err = proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Couldn't unmarshal")
		return
	}
}
