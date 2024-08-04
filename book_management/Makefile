start:
	GIN_MODE=release go run ./cmd/app/
	
dev:
	nodemon -e go --exec go run ./cmd/app/ --signal SIGTERM

mock:
	mockery --all