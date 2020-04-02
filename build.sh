#!/bin/bash   
for f in ./*
    do
	if [ -d $f ]; then
		cd $f
		go mod init ${f##*./}
		go build
		cd ..		
        fi   
done
