FROM golang:1.19.0-alpine AS base

RUN apk update && apk add git

WORKDIR /go/src
COPY ./backend .

FROM base AS prod

RUN go get -u github.com/cosmtrek/air && go build -o /go/bin/air github.com/cosmtrek/air \
&& go get github.com/gin-gonic/gin 

# air -c [tomlファイル名] // 設定ファイルを指定してair実行(WORKDIRに.air.tomlを配置しておくこと)
CMD ["air", "-c", ".air.toml"]

# CMD ["go", "run", "server.go"]