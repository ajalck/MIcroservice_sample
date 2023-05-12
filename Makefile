.PHONY : protos

protos:
	protoc --go_out=./usermgmt --go-grpc_out=./usermgmt usermgmt/usermgmt.proto

run_server:
	go run usermgmt_server/usermgmt_server.go

run_client:
	go run usermgmt_client/usermgmt_client.go