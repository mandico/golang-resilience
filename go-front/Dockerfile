# stage build
FROM golang:1.20.2 as build

WORKDIR /build

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/app

# stage imagem final
FROM alpine:3.17.3

RUN apk add tzdata \
    && cp /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime

WORKDIR /apps

COPY --from=build /build/app ./

CMD [ "./app" ]