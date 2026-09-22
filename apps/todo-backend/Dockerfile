FROM golang:1.27 AS build_image

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o b_app

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=build_image /app/b_app .

RUN groupadd -g 1000 deployuser && \
    useradd -m -u 1000 -g 1000 deployuser && \
    chown -R deployuser:deployuser /app

RUN touch todos.log && \
    chown -R deployuser:deployuser todos.log

USER deployuser

CMD ["./b_app"]