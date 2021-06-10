FROM golang:1.16.4 AS builder
LABEL n-guitar (https://github.com/n-guitar)
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:3.13 AS production
COPY --from=builder /app /app
WORKDIR /app
CMD [ "./main" ]