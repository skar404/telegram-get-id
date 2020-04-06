FROM golang:1.14.1-alpine3.11 as build

WORKDIR /app

COPY . .

RUN go build -o app && chmod +x app

FROM alpine:3.11

COPY --from=build /app .

EXPOSE 5000

CMD ["./app"]