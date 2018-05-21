#!/bin/bash

protoc -I protofiles/ --go_out=plugins=grpc:./src/product product.proto