mkdir build
GOOS="darwin" GOARCH="arm64" go build -o build/hasha-darwin-arm64/hasha -ldflags="-s -w" cmd/hasha/main.go
GOOS="linux" GOARCH="amd64" go build -o build/hasha-linux-amd64/hasha -ldflags="-s -w" cmd/hasha/main.go
GOOS="windows" GOARCH="amd64" go build -o build/hasha-windows-amd64/hasha.exe -ldflags="-s -w" cmd/hasha/main.go
