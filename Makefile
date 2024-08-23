ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Env vars used by goose app when running database migrations 
export GOOSE_DRIVER:=postgres
export GOOSE_DBSTRING:=$(DBSTRING)
export GOOSE_MIGRATION_DIR:=db/migrations

install:
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/air-verse/air@latest


templ:
	@templ generate -watch -proxy=http://localhost:8080/ -open-browser=false


tailwind:
	@npx tailwindcss -i assets/styles/my-styles.css -o assets/styles/styles.css --watch


air:
	@air


run: docker migrate/run
	make -j 3 templ tailwind air


test:
	@go test ./...


docker:
	@docker compose up db -d

	
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
	@templ generate
	@npx tailwindcss -i assets/styles/my-styles.css -o assets/styles/styles.css --minify
	@CGO_ENABLED=0; go build -o tmp/main cmd/webapp/main.go

