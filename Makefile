.PHONY : protos

protos:
	protoc --go_out=./usermgmt --go-grpc_out=./usermgmt usermgmt/usermgmt.proto

run:
	go run usermgmt_server/usermgmt_server.go