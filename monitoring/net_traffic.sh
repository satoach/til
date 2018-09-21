#!/bin/bash

readonly TARGET="eth0"

function bytes() {
    rcv=$(echo "$@" | cut -d ' ' -f 2)
    snd=$(echo "$@" | cut -d ' ' -f 10)

    echo "Receive,Transmit"
    echo "$rcv,$snd"
}

current=$(cat /proc/net/dev | grep ${TARGET} | tr -s ' ')
bytes $current

prevdata="traffic"
bytes $(cat ${prevdata}.*)

rm -f ${prevdata}.*
echo "$current" > ${prevdata}.$(date -Iseconds)
