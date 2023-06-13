FROM golang:1.20.0-alpine AS base

RUN apk update && apk add git

WORKDIR /go/src

COPY ./backend .


FROM base AS prod

RUN go get github.com/gin-gonic/gin \
    github.com/jinzhu/gorm \
    github.com/jinzhu/gorm/dialects/mysql \
    github.com/joho/godotenv

RUN go mod tidy

RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

RUN go install github.com/go-delve/delve/cmd/dlv@latest && \
    go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

CMD ["air", "-c", ".air.toml"]