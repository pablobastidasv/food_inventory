install:
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/air-verse/air@latest
	go mod vendor
	go mod tidy
	go mod download
	brew install tailwindcss


templ:
	@templ generate -watch -proxy=http://localhost:8080/


tailwind:
	@tailwindcss -i assets/styles/input.css -o assets/styles/output.css --watch


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
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING="user=postgres dbname=postgres sslmode=disable" GOOSE_MIGRATION_DIR="db/migrations" goose create $(name) sql


migrate/run:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING="host=localhost user=postgres dbname=postgres password=password sslmode=disable port=54321" GOOSE_MIGRATION_DIR="db/migrations" goose up


migrate/status:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING="host=localhost user=postgres dbname=postgres password=password sslmode=disable port=54321" GOOSE_MIGRATION_DIR="db/migrations" goose status


migrate/reset:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING="host=localhost user=postgres dbname=postgres password=password sslmode=disable port=54321" GOOSE_MIGRATION_DIR="db/migrations" goose reset

