build-win:
	GOOS=windows GOARCH=386 CGO_ENABLED=1 CXX=i686-w64-mingw32-g++ CC=i686-c64-mingw32-gcc go build -o mason-win.exe .
	cp mason-win.exe /mnt/c/Users/pete/Desktop/.

build:
	GOOS=linux GOARCH=amd64 go build -o mason .