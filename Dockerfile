FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o bin/reporter cmd/reporter/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/bin/reporter bin/reporter
ENTRYPOINT [ "./bin/reporter" ]