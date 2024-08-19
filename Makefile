install:
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest


templ:
	@templ generate


build: 
	@go build -o build/main cmd/webapp/main.go 


run: templ docker migrate/run
	@go run cmd/webapp/main.go 


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

