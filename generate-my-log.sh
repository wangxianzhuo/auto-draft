#!/bin/bash
currentFile="/home/shangjie/myLogs/current.log"
if [ -f "$currentFile" ]; then
  yesterday=`date -d yesterday +%Y-%m-%d`
  if [ -f "/home/shangjie/myLogs/$yesterday.log" ]; then
    mv /home/shangjie/myLogs/"$yesterday".log /home/shangjie/myLogs/"$yesterday".log.back
  fi
  mv /home/shangjie/myLogs/current.log /home/shangjie/myLogs/"$yesterday".log
fi
touch "$currentFile"
