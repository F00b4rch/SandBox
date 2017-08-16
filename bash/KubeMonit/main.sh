#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

#/ Usage: ./main.sh
#/ Description: This script monitore kubernetes infrastructure
#/ Examples: ./main.sh
#/ Options:
#/   --help: Display this help message
usage() { grep '^#/' "$0" | cut -c4- ; exit 0 ; }
expr "$*" : ".*--help" > /dev/null && usage

readonly LOG_FILE="/tmp/$(basename "$0").log"
info()    { echo "$(date -u) [INFO]    $*" | tee -a "$LOG_FILE" >&2 ; }
warning() { echo "$(date -u) [WARNING] $*" | tee -a "$LOG_FILE" >&2 ; }
error()   { echo "$(date -u) [ERROR]   $*" | tee -a "$LOG_FILE" >&2 ; }
fatal()   { echo "$(date -u) [FATAL]   $*" | tee -a "$LOG_FILE" >&2 ; exit 1 ; }

#/ sourcing vars:
source config.sh

if [[ "${BASH_SOURCE[0]}" = "$0" ]]; then

    #/ If root :
	if [[ $EUID -eq 0 ]]; then
    	echo "This script must be run as root"
    	exit 1
    fi


    while true; do


        #/ Func getting NODES
        getNodes() {
    		currentNodes=$(kubectl get nodes | grep -Ec '^gke')
    		if [[ "$currentNodes" = "$sumNodes" ]]; then
    			info "Number of nodes gets $currentNodes wanted $sumNodes"
    		else
    			warning "Number of nodes mismatch, wanted $sumNodes, have $currentNodes"
    		fi
    	}

        #/ Function call
    	getNodes

        #/ Func verifying STATUS
        nodeStatus() {
            nodeName=$(kubectl get nodes | grep -E '^gke' | awk '{print $1}')
            for i in $nodeName
            do status=$(kubectl get nodes $i | grep -E '^gke' | awk '{print $2}')
                if [ $status != "Ready" ]; then
                    echo "Fail"
                else
                    info "Nodes $i ready"
                fi
            done
        }

        nodeStatus

        sleep 2

    done

fi
