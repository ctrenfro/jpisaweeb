run: build
	@./bin/app

build:
	@go build -o bin/app cmd/app/main.go

css:
	./tailwindcss -i public/css/input.css -o public/css/output.css --watch

templ:
	templ generate --watch --proxy=http://localhost:3000

fmt: 
	templ fmt ./views .