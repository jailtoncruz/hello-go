FROM golang:1.21.3-alpine3.18 as build

WORKDIR /build

COPY . .

RUN [ "go", "build", "." ]

FROM alpine:3.14

WORKDIR /app

COPY --from=build /build/hello-world .

EXPOSE 3000

CMD ["./hello-world"]