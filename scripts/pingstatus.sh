#!/usr/bin/env bash

IP=$1
res=''
ping -c 1 -W 1 $IP &> /dev/null && res=success || res=fail
printf $res
