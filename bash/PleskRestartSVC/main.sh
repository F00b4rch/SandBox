#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

#/ Usage: ./main.sh
#/ Description: This script monitore Plesk Web services
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
# cleanup() {
#     rm "$LOG_FILE"
# }
# trap cleanup EXIT


# Functions 

    # Func mail
    sendMail() {
        echo "$(date) : $1 restarted on $HOSTNAME" | mail -s "$1 restarted on $HOSTNAME" "my@mail"
    }

    # Func nginx
    nginx() {
    if pgrep nginx > /dev/null
    then
#        info "Nginx service is up."
    :
    else
            warning "Nginx service down, restarting it !"
            if /etc/init.d/nginx restart > /dev/null
            then
                info "Nginx service restarted"
                sendMail "nginx"
            fi
    fi
    }

    # Func php-fpm
    php-fpm() {
    if pgrep php5-fpm > /dev/null
    then
#        info "Php-FPM service is up."
    :
    else
           warning "Php-FPM service down, restarting it !"
           if /etc/init.d/php5-fpm restart > /dev/null
           then
               info "Php-FPM service restarted"
               sendMail "php-FPM"
           fi
    fi
    }

    # Func mysql
    mysql(){
    if pgrep mysql > /dev/null
    then
#        info "Mysql service is up."
    :
    else
        warning "Mysql service down, restarting it !"
        if /etc/init.d/mysql restart > /dev/null 
        then
            info "Mysql service restarted"
            sendMail "mysql"
        fi
    fi
    }

    # Func apache
    apache(){
    if pgrep apache2 > /dev/null
    then
#        info "Apache2 service is up."
    :
    else
        warning "Apache2 service down, restarting it !"
        if /etc/init.d/apache2 restart > /dev/null 
        then
            info "Apache2 service restarted"
            sendMail "apache2"
        fi
    fi
    }

if [[ "${BASH_SOURCE[0]}" = "$0" ]]; then

    # If root :
	if [[ $EUID -ne 0 ]]; then
    	echo "This script must be run as root"
    	exit 1
    fi

    info "Starting..."

    while true; do

        # Functions call
        nginx    	
        sleep 5
        php-fpm
        sleep 5
        mysql
        sleep 5
        apache
        sleep 5

    done

fi
