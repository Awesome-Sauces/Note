#!/bin/bash
pkg="go"
which $pkg > /dev/null 2>&1
if [ $? == 0 ]
then
	read -p "Golang is not installed. Would you like to install it? y/n : " request
	if  [ $request == "y" ]
		then
			# Installs Golang 1.17
			(cd /tmp; wget https://golang.org/dl/go1.17.linux-amd64.tar.gz)
			(cd /tmp; tar -xzvf go1.17.linux-amd64.tar.gz)
			(cd /tmp; sudo mv go /usr/local)
			(cd /tmp; rm -rf go)
			echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
			echo 'export GOPATH=$HOME/go' >> ~/.bashrc
			echo 'export PATH=$GOPATH/bin:$GOROOT/bin:$PATH' >> ~/.bashrc
			source ~/.bashrc
fi
fi

# Grabs Newest version of Note
wget https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/note.go
wget https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/structs.go
wget https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/config.go
wget https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/commands.go
wget https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/go.mod
wget https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/go.sum
wget https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/build.sh
sh build.sh
sudo mv -i note /usr/bin
rm note.go
rm structs.go
rm config.go
rm commands.go
rm go.sum
rm go.mod
rm build.sh
sudo mkdir ~/note
wget https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/colorConfig.json
sudo mv -i colorConfig.json ~/note
