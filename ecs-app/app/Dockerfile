FROM golang:1.23-alpine AS builder
WORKDIR /build
ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV GOOS=linux 
ENV GOARCH=amd64

COPY . .
RUN go build -ldflags '-s -w' -o app

FROM scratch
COPY --from=builder /build/app ./app
EXPOSE 8080
ENTRYPOINT ["/app"]