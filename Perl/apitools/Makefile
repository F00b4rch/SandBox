SHELL := /bin/bash 

build:
		sudo docker build -t perlapitools .
		
runct:
		sudo docker run -it --rm --name running-perlapitools -p 8080:8080 -v "${PWD}"/:/usr/src/perl/ perlapitools /bin/bash start.sh

rmct:
		sudo docker stop running-perlapitools
		sudo docker rm running-perlapitools

rmi:
		sudo docker rmi perlapitools
