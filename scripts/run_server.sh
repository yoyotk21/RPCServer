#!/bin/bash

port="8000"
welcome_message="Welcome,"

while getopts p:m: flag
do 
    case "${flag}" in 
        p) port=${OPTARG};;
        m) welcome_message=${OPTARG};;
    esac
done 

echo "Port: $port"
echo "Welcome message: $welcome_message"

echo "Building..."
go build -o ../cmd/server ../cmd/server
echo "Done building!"


../cmd/server/server -port="${port}" -m="${welcome_message}" 
    
