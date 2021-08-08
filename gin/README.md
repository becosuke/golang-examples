```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' controller/hello/main.go
```
