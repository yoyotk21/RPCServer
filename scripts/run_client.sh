#!/bin/bash

# default values
name="human"
port="8000"
address="localhost"

while getopts n:p:a: flag 
do
    case "${flag}" in 
        n) name=${OPTARG};;
        p) port=${OPTARG};;
        a) address=${OPTARG};;
    esac
done

echo "Name: $name"
echo "Port: $port"
echo "Address: $address"

echo "Building..."
go build -o ../cmd/client  ../cmd/client
echo "Done building!"


../cmd/client/client \
    -name ${name} \
    -port ${port} \
    -addr ${address}