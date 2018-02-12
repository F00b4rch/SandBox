#!/bin/bash
# F00b4rch
#RHEL
clear
echo "updating cache and system"
yum makecache fast && yum update -y > /dev/null
echo "[OK]"
echo "installing tools"
yum install https://www.percona.com/redir/downloads/percona-release/redhat/0.0-1/percona-release-0.0-1.x86_64.rpm -y > /dev/null
yum install sysbench hdparm -y > /dev/null
echo "[OK]"
# CPU test
sleep 1
echo -ne "\n ###CPU TEST###\n"
sysbench --test=cpu --cpu-max-prime=20000 --num-threads=2 run 
echo -ne "\n ###THREADS TEST###\n"
# Threads test
sysbench --test=threads --thread-locks=1 --max-time=20 run
# IO test
echo -ne "\nIO TESTING\n"
df -h
echo -ne "\nEnter partition : \n" 
read partition 
# direct read
echo -ne "\n###Direct read###\n"
hdparm -t /dev/$partition
# cached read
echo -ne "\n###Cached read###\n"
hdparm -T /dev/$partition
# RAM test
echo -ne "\n###RAM TEST###\n"
dd if=/dev/zero of=/tmp/testfile bs=1M count=4000 
