#!/bin/bash

set -euo pipefail
IFS=$'\n\t'

#/ Usage: ./main.sh
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

# Debug
# Cleaning if exit
#cleanup() {
#    rm "$LOG_FILE"
#}
#trap cleanup EXIT

info "Starting running..."

while true ; do
    check=$(uptime | awk '{print $10}' | grep -Eo '^[0-9]')
    if [ "$check" -gt 8 ]; then
        service mysql restart
        info "MySQL service restarted"
        apachectl restart
        info "APACHE2 service restarted"
        echo "$(date) : MySQL/APACHE2 restarted on $HOSTNAME" | mail -s "MySQL/APACHE2 restarted on $HOSTNAME" "my@mail"
    fi
    sleep 40
done
