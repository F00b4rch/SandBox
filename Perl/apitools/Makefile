SHELL := /bin/bash 

start:
		sudo docker build -t perlapitools . && \
		sudo docker run --name running-perlapitools -p 8080:8080 perlapitools

clean:
		sudo docker stop running-perlapitools
		sudo docker rm running-perlapitools
		sudo docker rmi perlapitools
