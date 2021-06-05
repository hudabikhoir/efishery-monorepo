FROM golang:alpine as builder

RUN mkdir -p /efishery/backend 
ADD . /efishery/backend/
WORKDIR /efishery/backend 

ENV GO111MODULE=on

COPY go.mod go.sum ./

RUN export GOPROXY=https://proxy.golang.org \
    && go mod download

# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main app/server/main.go
RUN GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" \
    -o main \
    app/main.go

FROM alpine

RUN addgroup -S appgroup && adduser -S -D -H -h /app appuser -G appgroup

COPY --from=builder /efishery/backend/main /app/
COPY --from=builder /efishery/backend/config/ /app/config/

USER appuser

WORKDIR /app

STOPSIGNAL SIGINT

EXPOSE 1323

ENTRYPOINT ["./main"]