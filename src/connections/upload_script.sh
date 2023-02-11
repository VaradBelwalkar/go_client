#!/bin/bash

file_or_folder=$1
path=$2
containerpath=$3
port=$4
server=$5
keyForRemoteServer=$6

if [ "$file_or_folder" = "file" ]; then
  scp -i $keyForRemoteServer -P $port $path root@$server:$containerpath
elif [ "$file_or_folder" = "folder" ]; then
  scp -r -i $keyForRemoteServer -P $port $path root@$server:$containerpath
else
  echo "Error: Invalid argument. Expecting either 'file' or 'folder'."
fi
