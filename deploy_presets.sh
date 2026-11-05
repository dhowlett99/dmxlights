#!/bin/bash

files=`ls -d projects/.[!.]*`
for filename in ${files}
do
	echo "Creating Directory for ${filename}"
	mkdir -p dmxlights.app/Contents/Resources/${filename}
done

files=`ls -d projects/*.yaml`
for filename in ${files}
do
	echo "Creating Project Yaml for ${filename}"
	if [ -e ${filename} ]
	then
		cp -fr ${filename} dmxlights.app/Contents/Resources/${filename}
	fi
done

files=`ls -d projects/.[!.]*/*`
for filename in ${files}
do
	echo "Creating File for ${filename}"
	if [ -e ${filename} ]
	then
		cp -fr ${filename} dmxlights.app/Contents/Resources/${filename}
	fi
done