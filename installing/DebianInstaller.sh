#!/bin/bash
pkg="go"
u="$USER"
which $pkg > /dev/null 2>&1
if [ $? == 0 ]
then
ls
# Grabs Newest version of Note
wget https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/note.go
wget https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/go.mod
go build note.go
sudo mv note /usr/bin
else
    read -p "Golang is not installed. Would you like to install it? y/n : " request
if  [ $request == "y" ]
then
# Installs Golang 1.17
(cd /tmp; wget https://golang.org/dl/go1.17.linux-amd64.tar.gz)
(cd /tmp; tar -xzvf go1.17.linux-amd64.tar.gz)
(cd /tmp; sudo mv go /usr/local)
(cd /tmp; rm -rf go)
echo 'export GOROOT=/usr/local/go' >> /home/$u/.bashrc
echo 'export GOPATH=$HOME/go' >> /home/$u/.bashrc
echo 'export PATH=$GOPATH/bin:$GOROOT/bin:$PATH' >> /home/$u/.bashrc
source /home/$u/bashrc
# Grabs Newest version of Note
wget https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/note.go
wget https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/go.mod
go build note.go
sudo mv note /usr/bin
rm note.go
rm go.mod
fi
fi