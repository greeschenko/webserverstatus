#!/usr/bin/env bash

round(){
    echo $(printf %.$2f $(echo "scale=$2;(((10^$2)*$1)+0.5)/(10^$2)" | bc))
};

RES=""
TOTAL=$(round `free -m | grep Mem: | awk '{print($2)}'` 1)
FREE=$(round `free -m | grep Mem: | awk '{print($4)}'` 1)
FREEPERS=$(round `free -m | grep Mem: | awk '{print(100*$4/$2)}'` 1)
BUSYPERS=$(round `free -m | grep Mem: | awk '{print(100 - (100*$4/$2))}'` 1)

if [ $# -eq 0 ]; then
    RES="total/free (free%)  $TOTAL/$FREE ($FREEPERS%)"
else
    if [[ $1 == status ]]; then
        FREEPERS=`awk '{print $FREEPERS*100}' <<<"$FREEPERS"`
        if [[ $FREEPERS -lt 1000 ]]; then
            RES="FAIL"
        elif [[ $FREEPERS -lt 3000 ]]; then
            RES="WARN"
        else
            RES="OK"
        fi
    elif [[ $1 == graph ]]; then
        RES="${BUSYPERS}"
    elif [[ $1 == graphmax ]]; then
        RES=100
    fi
fi

printf "%s" "${RES}"
