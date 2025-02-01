// proto_gen.go
//go:generate protoc --proto_path=../../protos --go_out=../../../ --go-grpc_out=../../../ ../../protos/tempo.proto

package tempo