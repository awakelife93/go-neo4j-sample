FROM golang:alpine AS builder

# scratch 경량 이미지를 위한 환경변수 주입
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod go.sum main.go ./

RUN go mod download

COPY . ./

RUN go build -o main ./

WORKDIR /dist

RUN cp /build/main .

FROM scratch

COPY --from=builder /dist/main .

ENTRYPOINT ["/main"]