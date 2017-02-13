#!/bin/sh
# Install auto draft
option="${1}"
DIR=$(cd "$(dirname "$0")"; pwd)
case $option in
   --dir )    		if [ -f "${2}" ]; then
				
			else
				mkdir -p "${2}"
			fi
			echo "store draft in ${2}"
			mv generate-my-log.sh generate-my-log
              		;;
   --help | -h | * )	echo "--dir <where to store draft>"  
			;;
esac
