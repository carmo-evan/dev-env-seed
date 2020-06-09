#!/bin/sh
echo 'Checking environment...'
echo $ENV
if [ $ENV = production ]; then
    echo 'Spinning up production go server'
    go build main.go
    ./main
else
    echo 'Spinning up development go server'
    CompileDaemon --build="go build main.go" --command=./main
fi;