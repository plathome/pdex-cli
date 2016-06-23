# pdex-cli

The Go version of the pdexchange cli tool

*Grab Latest Release*

executable file attached

```
curl -L https://github.com/plathome/pdex-cli/releases/download/#{tag#}/pdex-cli.sh > /usr/local/bin/pdex-cli
```

```
chmod +x /usr/local/bin/pdex-cli
```


*Local Running*

Step 1:

```
go build
```

Step 2:

```
go run cli.go
```


*Docker Based Running*

# Build the go lang Docker

```
docker build -t docker/go .
```

# Running

```
docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app docker/go sh -c 'go run cli.go'
```

# Building

```
docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app docker/go sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o pdex-cli'
```

# Build for different OS Architecture

* MAC Build *
```
docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app docker/go sh -c 'env GOOS=darwin GOARCH=386 go build -v -o pdex-cli-darwin-386'
```

```
docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app docker/go sh -c 'env GOOS=darwin GOARCH=amd64 go build -v -o pdex-cli-darwin-amd64'
```

* Linux Build *

```
docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app docker/go sh -c 'env GOOS=linux GOARCH=386 go build -v -o pdex-cli-linux-386'
```

```
docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app docker/go sh -c 'env GOOS=linux GOARCH=amd64 go build -v -o pdex-cli-linux-amd64'
```

```
docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app docker/go sh -c 'env GOOS=linux GOARCH=arm go build -v -o pdex-cli-linux-arm'
```

```
docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app docker/go sh -c 'env GOOS=linux GOARCH=arm64 go build -v -o pdex-cli-linux-arm64'
```

* Windows Build *

```
docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app docker/go sh -c 'env GOOS=windows GOARCH=386 go build -v -o pdex-cli-windows-386.exe'
```

```
docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app docker/go sh -c 'env GOOS=windows GOARCH=amd64 go build -v -o pdex-cli-windows-amd64.exe'
```
