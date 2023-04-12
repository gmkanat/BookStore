FROM golang:1.20-alpine AS build
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /book-store

FROM scratch
COPY --from=build /book-store /book-store
EXPOSE 8080
ENTRYPOINT ["/book-store"]
