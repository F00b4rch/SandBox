#!/bin/bash -
#===============================================================================
#
#          FILE: installdockerpcicentos.sh
#
#         USAGE: ./installdockerpcicentos.sh
#
#   DESCRIPTION:
#
#       OPTIONS: ---
#  REQUIREMENTS: ---
#          BUGS: ---
#         NOTES: ---
#        AUTHOR: Corentin Deret (DevOps Engineer)
#       CREATED: 10/07/2017 09:26:15
#      REVISION:  ---
#===============================================================================

yum makecache fast
yum update -y
yum install vim -y

yum remove -y docker \
              docker-common \
              docker-selinux \
              docker-engine

yum install -y yum-utils device-mapper-persistent-data lvm2

yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo

yum makecache fast

yum install docker-ce -y && systemctl start docker && systemctl enable docker && echo -ne '\n\n\n[OK]\n\n\n'
