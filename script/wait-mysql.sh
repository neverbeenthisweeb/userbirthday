#!/bin/sh
# Docker-compose check if mysql connection is ready
# https://stackoverflow.com/a/51641089

maxcounter=30
counter=1
while ! mysql -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" -h"0.0.0.0" -P"3306" -e "SHOW DATABASES;" > /dev/null 2>&1; do
    sleep 1
    counter=`expr $counter + 1`
    if [ $counter -gt $maxcounter ]; then
        >&2 echo "Too long waiting for MySQL; Failing."
        exit 1
    fi;
done