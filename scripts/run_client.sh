#!/bin/bash

# default values
index=0
port1="8000"
port2="8080"
address="localhost"

while getopts i:p1:p2:a: flag 
do
    case "${flag}" in 
        i) index=${OPTARG};;
        p1) port1=${OPTARG};;
        p2) port2=${OPTARG};;
        a) address=${OPTARG};;
    esac
done

echo "Index: $index"
echo "Port 1: $port1"
echo "Port 2: $port2"
echo "Address: $address"

echo "Building..."
go build -o ../cmd/client  ../cmd/client
echo "Done building!"


../cmd/client/client \
    -i ${index} \
    -p1 ${port1} \
    -p2 ${port2} \
    -a ${address}