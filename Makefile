test-all:
	go test .
	go test client/client_test.go
	go test server/server_test.go
