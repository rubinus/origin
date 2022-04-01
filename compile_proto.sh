#!/bin/bash

protoc --proto_path=proto --gofast_out=plugins=grpc:pb helloworld/helloworld.proto

protoc --proto_path=proto --gofast_out=plugins=grpc:pb weather/weather.proto

#在这里添加其它的proto
