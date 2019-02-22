CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o myapp lambda.go
zip lambda.zip myapp
