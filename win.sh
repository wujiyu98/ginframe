set GOROOT=D:\WorkSpace\Go
rice clean
rice embed-go -i ./router
set GOROOT=C:\Go
go build -o ./bin/front.exe ./cmd