FROM golang:1.14.1-alpine3.11 as build

WORKDIR /src

COPY . .

RUN go build -o app && chmod +x app

FROM alpine:3.11

COPY --from=build /src/app .

CMD ["./app"]