#!/bin/bash
#// To run this file. us the following command
#// fswatch -v . | xargs -n1 ./build.sh
#// ------------------------------------------
rm ../build/main
pwd
echo "------Building the Executable from------"
go build  -o ../build/main main.go
echo "----------------------------------------"
../build/main
echo "========================================"

echo "Built End"
