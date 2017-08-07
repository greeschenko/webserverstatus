#!/usr/bin/env bash

RES=''
LOGFILE=$1
INTERVAL=$2
COUNT=`wc -l $LOGFILE | awk '{print($1)}'`
TMPFILE=`echo -n $LOGFILE | md5sum | awk '{print($1)}'`
TMPFILE=/tmp/$TMPFILE

if [[ -f $TMPFILE ]]; then
    OLD=`cat $TMPFILE`
else
    OLD=$COUNT
fi

echo "${COUNT}" > $TMPFILE

RES=`expr $COUNT - $OLD`
RES=`expr $RES / $INTERVAL`

if [ $# -eq 2 ]; then
    RES="Load: ${RES} req/second"
else
    if [[ $3 == status ]]; then
        if [[ $RES -gt 500 ]]; then
            RES="FAIL"
        elif [[ $RES -gt 200 ]]; then
            RES="WARN"
        else
            RES="OK"
        fi
    elif [[ $3 == graph1 ]]; then
        RES="${RES}"
    elif [[ $3 == graphmax ]]; then
        RES=500
    fi
fi

printf "${RES}"
