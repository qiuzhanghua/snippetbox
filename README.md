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

```bash
go test -v -run="^TestPing$" ./cmd/web/
go test -v -run="^TestHumanDate$/^UTC$" ./cmd/web/
go test -v -skip="^TestHumanDate$" ./cmd/web/

# avoid cache
 go test -count=1 ./cmd/web
# clean cache
go clean -testcache
# fast failure
go test -failfast ./cmd/web
# race detector
go test -race ./cmd/web
```