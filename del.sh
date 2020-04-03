#!/bin/bash   
for f in ./*
    do
	if [ -d $f ]; then
		cd $f
		if [ -f "go.mod" ]; then 
    			rm go.mod
		fi
		if [ -x $f ]; then
			rm $f
		fi
		cd ..		
        fi   
done
