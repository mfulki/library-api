start:
	GIN_MODE=release go run ./cmd/app/
	
dev:
	nodemon -e go --exec go run ./cmd/app/ --signal SIGTERM

mock:
	mockery --all

protobuf: 
	 protoc --proto_path=pkg/protos pkg/protos/*.proto --go_out=./ --go-grpc_out=./            
