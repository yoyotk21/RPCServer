#!/bin/bash

port="8000"

while getopts p: flag
do 
    case "${flag}" in 
        p) port=${OPTARG};;
    esac
done 

echo "Port: $port"

echo "Building..."
go build -o ../cmd/server ../cmd/server
echo "Done building!"


../cmd/server/server -p="${port}"
    
