FROM golang:latest as builder

COPY ./src/main.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /soma

FROM scratch

COPY --from=builder /soma /soma

ENTRYPOINT ["/soma"]