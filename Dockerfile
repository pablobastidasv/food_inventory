FROM golang:alpine AS build_stage

WORKDIR /build

RUN apk add --no-cache make
RUN apk add --update nodejs npm

COPY go.mod go.mod 
COPY package.json package.json
COPY Makefile Makefile

RUN go mod download
RUN npm install
RUN make install

COPY . .

RUN make build


#Deploy Stage
FROM alpine:latest

WORKDIR /app

COPY --from=build_stage /build/tmp/main main
COPY assets assets

EXPOSE 8080

# USER nonroot:nonroot

ENTRYPOINT ["./main"]
