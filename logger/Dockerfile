FROM golang:alpine as builder

WORKDIR /go/src/logger
COPY . .

RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

RUN go get -d -v ./...
RUN go install -v ./...

# compile static executable
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/logger .


# step 2. build minified image

FROM scratch

# Import ca-certs from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy our static executable
COPY --from=builder /go/bin/logger /go/bin/logger

ENTRYPOINT ["/go/bin/logger"]
