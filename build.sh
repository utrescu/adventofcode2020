#!/bin/bash
i=1
for dir in $(ls -d day*)
do
    pushd "$dir" > /dev/null
    go build advent$i.go
    rm advent$i
    popd > /dev/null
    i=$(($i+1))
done
