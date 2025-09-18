#!/bin/bash

set -e

gcc -o cbot bot.c -lcjson
cp /lib/x86_64-linux-gnu/libcjson.so.1 . 
docker build -t c_example .
rm libcjson.so.1
