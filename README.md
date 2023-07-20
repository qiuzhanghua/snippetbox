# Learn with
Let's go!


For TLS:

```bash
mkdir tls
cd tls
go run $GOROOT/src/crypto/tls/generate_cert.go --rsa-bits=2048 -host=localhost
```

### Test
    
    ```bash
    go test -v ./...
    # or
    go test ./cmd/web
    ```
