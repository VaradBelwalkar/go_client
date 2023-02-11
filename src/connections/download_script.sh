#!/bin/bash

file_or_folder=$1
containerpath=$2
localpath=$3
port=$4
server=$5
keyForRemoteServer=$6
if [ "$file_or_folder" = "file" ]; then
  scp -i $keyForRemoteServer -P $port  root@$server:$containerpath $localpath
elif [ "$file_or_folder" = "folder" ]; then
  scp -r -i $keyForRemoteServer -P $port  root@$server:$containerpath $localpath
else
  echo "Error: Invalid argument. Expecting either 'file' or 'folder'."
fi
