#!/bin/bash

# Usage: slackpost "<webhook_url>" "<channel>" "<message>"

# Add your webhook here
webhook_url=""
if [[ $webhook_url == "" ]]
then
        echo "No webhook_url specified"
        exit 1
fi

shift

#add your chennel 
channel=""
if [[ $channel == "" ]]
then
        echo "No channel specified"
        exit 1
fi

shift

hostname=$(hostname)
docker_id=$(docker ps -aq)

while true ; do 

if [ $(docker ps -q | wc -l | grep 2) ] ; then
	:
else
	text="ALERT $hostname, there is no ct running."
escapedText=$(echo $text | sed 's/"/\"/g' | sed "s/'/\'/g" )
json="{\"channel\": \"$channel\", \"text\": \"$escapedText\"}"

curl -s -d "payload=$json" "$webhook_url"


fi

isup=$(docker inspect -f '{{.State.Running}}' $(docker ps -q))
res="true"

for i in $isup ;do

if [ $i == $res ] ; then
	:
else
        text="ALERT $hostname, CT state running false."
escapedText=$(echo $text | sed 's/"/\"/g' | sed "s/'/\'/g" )
json="{\"channel\": \"$channel\", \"text\": \"$escapedText\"}"

curl -s -d "payload=$json" "$webhook_url"

fi

done


sleep 5 ;

done
