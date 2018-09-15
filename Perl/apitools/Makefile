SHELL := /bin/bash 

build:
		sudo docker build -t perlapitools . && \
		sudo docker run --name running-perlapitools -p 8080:8080 perlapitools

delete:
		sudo docker rm running-perlapitools