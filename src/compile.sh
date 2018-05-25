echo "compiling on MAC"
pwd
go tool compile -I /Users/geoffrey/go/pkg/darwin_amd64 -o ../build/main.o main.go
echo "Compiled"