
PROTO_DIR = proto
PROTO_OUT = protoOut
PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')

generate:
	protoc -I${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. ${PROTO_DIR}/*.proto

clean:
	rm ${PROTO_OUT}/*.pb.go

build:
	go build .
