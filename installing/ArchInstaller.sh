#!/bin/bash
echo "Installing Note"
x=`pacman -Qs go`
if [ -n "$x" ]
then
    touch note.go
    touch go.mod
    // Grabs Newest version of Note
    curl https://raw.githubusercontent.com/Awesome-Sauces/Note/Note-v1.0.0-BETA/src/note.go > note.go
    curl https://raw.githubusercontent.com/Awesome-Sauces/Note/Note-v1.0.0-BETA/src/go.mod > go.mod
    go build note.go
    sudo mv note /usr/bin
    rm note.go
    rm go.mod
else
    sudo pacman -S go
    touch note.go
    touch go.mod
    // Grabs Newest version of Note
    curl https://raw.githubusercontent.com/Awesome-Sauces/Note/Note-v1.0.0-BETA/src/note.go > note.go
    curl https://raw.githubusercontent.com/Awesome-Sauces/Note/Note-v1.0.0-BETA/src/go.mod > go.mod
    go build note.go
    sudo mv note /usr/bin
    rm note.go
    rm go.mod
fi