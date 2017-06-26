#!/bin/bash - 
#===============================================================================
#
#          FILE: pars.sh
# 
#         USAGE: ./pars.sh file
# 
#   DESCRIPTION: 
# 
#       OPTIONS: ---
#  REQUIREMENTS: ---
#          BUGS: ---
#         NOTES: ---
#        AUTHOR: Corentin Deret (Administrateur Système DevOps)
#       CREATED: 26/06/2017 10:38:54
#      REVISION:  ---
#===============================================================================

file=$1

# nombre total de requete
totalReq=$(cat $file | wc -l)
echo -ne "\ntotal request = $totalReq"

# nombre total ip
totalIp=$(grep -Eo ':[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\ ' $file | wc -l)
echo -ne "\ntotal ip = $totalIp"

cat $file | grep -Eo ':[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\ ' > ip.txt
sed -i 's/://g' ip.txt
echo -ne '\n\n****Top 10 IP****\n\n'
sort ip.txt | uniq -c | sort -n -r | head -10
rm ip.txt

echo -ne '\n\n****Top 10 requetes*****\n\n'
cat $file | awk '{print $7}' > requests
sort requests | uniq -c | sort -n -r | head -10

rm requests

echo -ne'\n\n***** REQUETES /h *****\n\n'
for i in $(seq 0 9) ; do echo -ne "$i heures\n" ; grep "2017:0$i:" $file | wc -l ; done && for i in $(seq 10 23) ; do echo -ne "$i heures\n" ; grep "2017:$i:" $file | wc -l ; done 
