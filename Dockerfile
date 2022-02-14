FROM golang:1.17 as builder

WORKDIR /build

ADD ./ /build/

RUN CGO_ENABLED=0 GOOS=linux go build -a -o hello-world .

FROM alpine:3.15

COPY --from=builder /build/hello-world .

CMD [ "./hello-world" ]