test:
	go test ./...
watch:
	wgo go run ./cmd/wgserver/main.go -host=3000
