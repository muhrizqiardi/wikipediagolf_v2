test:
	go test ./...
watch:
	wgo -xdir=./client/dist bun run dev :: go run ./cmd/wgserver/main.go -port=3001 -migrate=true 
