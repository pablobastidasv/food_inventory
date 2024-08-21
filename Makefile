include .env

install:
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/air-verse/air@latest
	go mod vendor
	go mod tidy
	go mod download
	npm install


templ:
	@templ generate -watch -proxy=http://localhost:8080/ -open-browser=false


tailwind:
	@npx tailwindcss -i assets/styles/input.css -o assets/styles/output.css --watch


air:
	@air


build: 
	@go build -o build/main cmd/webapp/main.go 


run: docker migrate/run
	make -j 3 templ tailwind air


test:
	@go test ./...


docker:
	@docker compose up -d

	
migrate/create: 
	@goose create $(name) sql


migrate/run:
	@goose up


migrate/status:
	@goose status


migrate/reset:
	@goose reset


# apk add --no-cache make
# apk add --update nodejs npm
build:
	@make install
	@templ generate
	@npx tailwindcss -i assets/styles/input.css -o assets/styles/output.css --minify
	@CGO_ENABLED=0; go build -o tmp/main cmd/webapp/main.go

