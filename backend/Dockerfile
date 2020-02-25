FROM golang as builder

WORKDIR /usr/app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o product-compare .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

ENV PATH /root

WORKDIR /root/

COPY --from=builder /usr/app/product-compare .

EXPOSE 8080

CMD ["./product-compare"] 