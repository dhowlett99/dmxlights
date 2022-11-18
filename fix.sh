#!/bin/sh

sed 's/\<\/dict\>/    \<key\>NSMicrophoneUsageDescription\<\/key\>\
    \<string\>Required for music trigger\<\/string\>\
\<\/dict\>/' $1

