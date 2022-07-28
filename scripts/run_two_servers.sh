#!/bin/bash

port1="8000"
port2="8080"

while getopts p1:p2: flag
do 
    case "${flag}" in 
        p1) port1=${OPTARG};;
        p2) port2=${OPTARG};;
    esac
done

echo "Port 1: $port1"
echo "Port 2: $port2"

echo "Building..."
go build -o ../cmd/server ../cmd/server
echo "Done building!"

(trap "kill 0" SIGINT; ../cmd/server/server -p="${port1}" & ../cmd/server/server -p="${port2}")