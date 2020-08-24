FROM golang:1.14.4-alpine AS compiler

WORKDIR /app/src

COPY . .

RUN go build -o gocandela

FROM alpine

COPY --from=compiler /app/src/gocandela /gocandela

ENTRYPOINT ["/gocandela"]
