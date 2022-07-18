#!/bin/bash
pkg="go"
which $pkg > /dev/null 2>&1
if [ $? != 0 ]
then
    read -p "Golang is not installed. Would you like to install it? y/n " request
	if  [ $request == "y" ]
	then
   	 sudo pacman -S go
fi
fi

# Grabs Newest version of Note
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/note.go --output note.go
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/structs.go --output structs.go
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/config.go --output config.go
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/commands.go --output commands.go
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/go.mod --output go.mod
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/go.sum --output go.sum
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/build.sh --output build.sh
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
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/colorConfig.json > colorConfig.json
sudo mv -i colorConfig.json ~/note
