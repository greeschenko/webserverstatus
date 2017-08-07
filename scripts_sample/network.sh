#!/usr/bin/env bash

RES=""
INTF=$1
INTERVAL=$2

R1=`cat /sys/class/net/$INTF/statistics/rx_bytes`
T1=`cat /sys/class/net/$INTF/statistics/tx_bytes`
ER1=`cat /sys/class/net/$INTF/statistics/rx_errors`
ET1=`cat /sys/class/net/$INTF/statistics/tx_errors`
DR1=`cat /sys/class/net/$INTF/statistics/rx_dropped`
DT1=`cat /sys/class/net/$INTF/statistics/tx_dropped`
LR1=0
LT1=0


if [[ -f /tmp/wstat_last_rx ]]; then
    LR1=`cat /tmp/wstat_last_rx`
else
    echo "${R1}" > /tmp/wstat_last_rx
    LR1=$R1
fi

if [[ -f /tmp/wstat_last_tx ]]; then
    LT1=`cat /tmp/wstat_last_tx`
else
    echo "${T1}" > /tmp/wstat_last_tx
    LT1=$T1
fi

echo "${R1}" > /tmp/wstat_last_rx
echo "${T1}" > /tmp/wstat_last_tx

TBPS=`expr $T1 - $LT1`
TBPS=`expr $TBPS / $INTERVAL`
TKBPS=`expr $TBPS / 1024`
RBPS=`expr $R1 - $LR1`
RBPS=`expr $RBPS / $INTERVAL`
RKBPS=`expr $RBPS / 1024`
RALL=`expr $R1 / 1024`
TALL=`expr $T1 / 1024`

if [ $# -eq 2 ]; then
    RES="TX $INTF: $TALL kB ($TKBPS kB/s) errors/drops $ET1/$DT1 RX $INTF: $RALL kB ($RKBPS kB/s) errors/drops $ER1/$DR1"
else
    if [[ $3 == status ]]; then
        RES="OK"

        if [[ $TKBPS -gt 5000 ]]; then
            RES="FAIL"
        elif [[ $TKBPS -gt 3000 ]]; then
            RES="WARN"
        fi

        if [[ $RKBPS -gt 5000 ]]; then
            RES="FAIL"
        elif [[ $RKBPS -gt 3000 ]]; then
            RES="WARN"
        fi

        if [[ $ER1 -gt 100 ]]; then
            RES="FAIL"
        elif [[ $ER1 -gt 0 ]]; then
            RES="WARN"
        fi
    elif [[ $3 == graph1 ]]; then
        RES="${TKBPS}"
    elif [[ $3 == graph2 ]]; then
        RES="${RKBPS}"
    elif [[ $3 == graphmax ]]; then
        RES=5000
    fi
fi

printf "%s" "${RES}"
