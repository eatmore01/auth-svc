#!/bin/bash


LOGS_FILE_NAME="$(date +%Y-%m:%d-%H:%M:%S)-logs.log"


while true; do
    FORMATED_DATE=$(date +%Y/%m:%d/%H:%M:%S)
    if curl -f http://0.0.0.0:9091; then
       echo "$FORMATED_DATE: Server is up" >> ./logs/$LOGS_FILE_NAME
    else 
        echo "$FORMATED_DATE: Server down" >> ./logs/$LOGS_FILE_NAME
        exit 1
    fi
    sleep 5
done