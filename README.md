# pdex-cli

pdex : The Go version of the pdexchange cli tool

*Grab Latest Release*

executable file attached

```
curl -L https://github.com/plathome/pdex-cli/releases/download/#{tag#}/pdex-`uname -s`-`uname -m` > /usr/local/bin/pdex
```

```
chmod +x /usr/local/bin/pdex
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
docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app docker/go sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o pdex'
```

# Build for different OS Architecture

*MAC Build*
```
docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app docker/go sh -c 'env GOOS=darwin GOARCH=386 go build -v -o pdex-Darwin-x86'
```

```
docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app docker/go sh -c 'env GOOS=darwin GOARCH=amd64 go build -v -o pdex-Darwin-x86_64'
```

*Linux Build*

```
docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app docker/go sh -c 'env GOOS=linux GOARCH=386 go build -v -o pdex-Linux-x86'
```

```
docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app docker/go sh -c 'env GOOS=linux GOARCH=amd64 go build -v -o pdex-Linux-x86_64'
```

```
docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app docker/go sh -c 'env GOOS=linux GOARCH=arm go build -v -o pdex-Linux-arm'
```

```
docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app docker/go sh -c 'env GOOS=linux GOARCH=arm64 go build -v -o pdex-Linux-arm64'
```

*Windows Build*

```
docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app docker/go sh -c 'env GOOS=windows GOARCH=386 go build -v -o pdex-Windows-x86.exe'
```

```
docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app docker/go sh -c 'env GOOS=windows GOARCH=amd64 go build -v -o pdex-Windows-x86_64.exe'
```
