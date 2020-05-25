FROM golang:1.13-buster as builder
MAINTAINER "Gian Marco Mennecozzi"
WORKDIR /haaukins

COPY . .
RUN go build -o server .

FROM gcr.io/distroless/base-debian10
COPY --from=builder /haaukins /
EXPOSE 50051
CMD ["/server"]
