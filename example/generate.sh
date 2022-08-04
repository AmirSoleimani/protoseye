#! /bin/bash

find . -name '*.proto' -exec protoc -I=. \
    --protoseye_out=./outputs {} \;
