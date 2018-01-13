#! /bin/bash

protoc -I=proto --go_out=plugins=grpc:pb  proto/courier.proto