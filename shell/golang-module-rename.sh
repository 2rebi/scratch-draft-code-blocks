#!/bin/bash

OLD_MODULE=`head -n 1 go.mod | awk '{print $2}'`;

read -p "$OLD_MODULE, rename module name : " NEW_MODULE;
echo $NEW_MODULE;

go mod edit -module $NEW_MODULE;

for FILE in `find . -type f -name '*.go'`
do
  TEXT=`cat $FILE`;
  echo "${TEXT//$OLD_MODULE/$NEW_MODULE}" >| $FILE;
done;
