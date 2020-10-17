#!/bin/bash
# Generate Micro pb files.
echo "generating protoc for project $NAME"

if [ -z "$GOPATH" ]  
then  
    export GOPATH=~/go
else  
    echo $GOPATH/pkg/mod
fi
protoc \
-I . \
-I $GOPATH/src/github.com/envoyproxy/protoc-gen-validate \
-I $GOPATH/src \
--micro_out=. \
--go_out=. \
--validate_out="lang=go:." \
$NAME/message.proto

echo $NAME was generated!
