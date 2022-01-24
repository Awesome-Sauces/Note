#!/bin/bash
pkg="go"
which $pkg > /dev/null 2>&1
if [ $? == 0 ]
then
    touch note.go
    touch go.mod
    // Grabs Newest version of Note
    curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/note.go > note.go
    curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/go.mod > go.mod
    go build note.go
    sudo mv note /usr/bin
    rm note.go
    rm go.mod
else
    read -p "Golang is not installed. Would you like to install it? y/n " request
if  [ $request == "y" ]
then
    sudo pacman -S go
    touch note.go
    touch go.mod
    // Grabs Newest version of Note
    curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/note.go > note.go
    curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/go.mod > go.mod
    go build note.go
    sudo mv note /usr/bin
    rm note.go
    rm go.mod
fi
fi