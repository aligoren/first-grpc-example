create:
	protoc --proto_path=proto proto/*.proto --go_out=gen/
	protoc --proto_path=proto proto/*.proto --go-grpc_out=gen/

run-server:
	go build -o ${CURDIR}/server/server.exe ${CURDIR}/server/server.go
	${CURDIR}/server/server.exe

run-client:
	go build -o ${CURDIR}/client/client.exe ${CURDIR}/client/client.go
	${CURDIR}/client/client.exe