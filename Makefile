test:
	go test ./...
watch:
	wgo npm run build :: go run ./cmd/wgserver/main.go -host=3000  
