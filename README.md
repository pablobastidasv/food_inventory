# Bastriguez Inventory

## Used libraries

### Prod

- https://github.com/a-h/templ 
- https://github.com/google/uuid 
- https://github.com/joho/godotenv 
- https://github.com/labstack/echo/v4 
- https://github.com/samber/slog-echo 
- https://github.com/markbates/goth 
- https://github.com/lib/pq 

### Testing

- https://github.com/stretchr/testify 
- https://github.com/tinygg/gofaker 


## Running the app locally

Define in your system the variables `AUTH_CLIENT_ID` and `AUTH_CLIENT_SECRET` then use `make` to run the full app

### Dev mode

```bash
make run/dev
```

### Via Docker

```bash
make run/docker
```

