test:
	go test ./...
watch:
	wgo -xdir=internal/view/asset/dist npm run build :: go run ./cmd/wgserver/main.go -port=3001 -migrate=true 
