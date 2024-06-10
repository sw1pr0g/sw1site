build:
	@GOARCH=wasm GOOS=js go build -o web/app.wasm ./cmd/app
	@go build -o web/sw1site ./cmd/app

clean:
	@go clean ./...
	@-rm web/sw1site