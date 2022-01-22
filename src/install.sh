#!bin/bash
echo "Installing Note"
x=`pacman -Qs go`
if [ -n "$x" ]
 then 
    echo "Hello Buddy"
 else 
    echo "Hello"
fi


#!/bin/bash
echo "Installing Note"
x=`pacman -Qs go`
if [ -n "$x" ]
then
    go build note.go
    sudo cp note /usr/bin
else
    sudo pacman -S go
    go build note.go
    sudo cp note /usr/bin
fi