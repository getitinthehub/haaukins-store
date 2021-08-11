FROM golang:1.16.5-buster as builder
MAINTAINER "Gian Marco Mennecozzi"
WORKDIR /haaukins

COPY . .
RUN go mod download
RUN go build -o server .

FROM gcr.io/distroless/base-debian10
COPY --from=builder /haaukins /
EXPOSE 50051
CMD ["/server"]
