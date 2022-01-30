#!/bin/bash

setup:
	cp -n .env.dist .env || true
	go build -o main .

run:
	./main