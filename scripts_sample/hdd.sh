#!/usr/bin/env bash

DSK=/dev/sda
RES=""
USE=`df -H | grep ${DSK} | awk '{print($5)}'`
USE="${USE/\%/}"
if [ $# -eq 0 ]; then
    RES=`df -H | grep ${DSK} | awk '{print("size/used/used%",$2,$3,$5)}'`
else
    if [[ $1 == status ]]; then
        if [[ $USE -gt 95 ]]; then
            RES="FAIL"
        elif [[ $USE -gt 75 ]]; then
            RES="WARN"
        else
            RES="OK"
        fi
    elif [[ $1 == graph ]]; then
        RES="${USE}"
    elif [[ $1 == graphmax ]]; then
        RES=100
    fi
fi
printf "%s" "${RES}";
