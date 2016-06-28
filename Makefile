output=pdex

clean:
	rm -f ${output}-*

all: clean linux linux64 linuxarm linuxarm64 mac64 mac windows windows64

linux:
	GOOS=linux GOARCH=386 go build -v -o ${output}-Linux-x86 cli.go

linux64:
	GOOS=linux GOARCH=amd64 go build -v -o ${output}-Linux-x86_64 cli.go

linuxarm:
	GOOS=linux GOARCH=arm go build -v -o ${output}-Linux-arm cli.go

linuxarm64:
	GOOS=linux GOARCH=arm64 go build -v -o ${output}-Linux-arm64 cli.go

mac64:
	GOOS=darwin GOARCH=amd64 go build -v -o ${output}-Darwin-x86_64 cli.go

mac:
	GOOS=darwin GOARCH=386 go build -v -o ${output}-Darwin-x86 cli.go

windows:
	GOOS=windows GOARCH=386 go build -v -o ${output}-Windows-x86.exe cli.go

windows64:
	GOOS=windows GOARCH=amd64 go build -v -o ${output}-Windows-x86_64.exe cli.go
