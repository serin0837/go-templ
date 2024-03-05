run:
	templ generate
	go run cmd/main.go

hot-reload:
	templ generate --watch --cmd="go run cmd/main.go" --proxy="http://localhost:3000"