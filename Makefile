test:
	go test ./...
watch:
	wgo -file=.go go run ./cmd/wgserver/main.go -host=3000
