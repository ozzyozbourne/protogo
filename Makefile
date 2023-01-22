
PROTO_DIR = proto
PROTO_OUT = protoOut
PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')

generate:
	protoc -I${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. ${PROTO_DIR}/*.proto

clean_proto:
	rm ${PROTO_OUT}/*.pb.go

clean_output:
	rm grpcgo

build:
	go build .

run:
	./grpcgo

all: clean_output clean_proto generate build run