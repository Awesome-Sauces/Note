#!/bin/bash
pkg="go"
which $pkg > /dev/null 2>&1
if [ $? == 0 ]
then
ls
# Grabs Newest version of Note
wget https://raw.githubusercontent.com/Awesome-Sauces/Note/Note-v1.0.0-BETA/src/note.go
wget https://raw.githubusercontent.com/Awesome-Sauces/Note/Note-v1.0.0-BETA/src/go.mod
go build note.go
sudo cp note /usr/bin
rm note.go
rm go.mod
else
    read -p "$pkg is not installed. Would you like to install it? y/n " request
if  [ $request == "y" ]
then
# Installs Golang 1.17
wget https://golang.org/dl/go1.17.6.linux-amd64.tar.gz
tar -xzvf go1.17.6.linux-amd64.tar.gz
sudo mv go /usr/local
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
rm go1.17.6.linux-amd64.tar.gz
# Grabs Newest version of Note
wget https://raw.githubusercontent.com/Awesome-Sauces/Note/Note-v1.0.0-BETA/src/note.go
wget https://raw.githubusercontent.com/Awesome-Sauces/Note/Note-v1.0.0-BETA/src/go.mod
go build note.go
sudo cp note /usr/bin
rm note.go
rm go.mod
fi
fi
