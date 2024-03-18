FROM golang:alpine as builder
LABEL maintainer="Alan Kelly <alan.kelly.london@gmail.com>"
COPY . /build
WORKDIR /build
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o htmx

######## Start a new stage from scratch #######
FROM alpine:latest  
WORKDIR /
COPY --from=builder /build .
EXPOSE 10000
CMD ["/htmx"]