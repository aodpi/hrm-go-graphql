FROM golang:alpine AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY ./src .
RUN go build -o main .

FROM alpine AS final
WORKDIR /app
ENV SERVER_ADDRESS=0.0.0.0:4000
EXPOSE 4000
COPY --from=build /src/main /src/config/dev.yaml  /app/
CMD ["/app/main"]
