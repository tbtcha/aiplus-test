FROM golang:latest AS build
WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 go build -o /bin/server .

FROM alpine:latest AS final

COPY --from=build /bin/server /bin/

COPY .env .

EXPOSE 8080

ENTRYPOINT [ "/bin/server" ]
