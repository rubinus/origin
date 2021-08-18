#!/bin/bash

protoc --proto_path=proto --gofast_out=plugins=grpc:pb helloworld/helloworld.proto

#在这里添加其它的proto