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

sudo mkdir ~/note
touch theme.rocky
sudo mv -i themes.rocky ~/note

touch note.go
touch structs.go
touch config.go
touch commands.go
touch build.sh
touch go.sum
touch go.mod
touch TextEditor.go
# Grabs Newest version of Note
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/note.go > note.go
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/structs.go > structs.go
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/config.go > config.go
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/commands.go > commands.go
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/go.mod > go.mod
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/TextEditor.go > TextEditor.go
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/go.sum > go.sum
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/build.sh > build.sh
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/themes/theme.rocky > theme.rocky
sh build.sh
sudo mv -i note /usr/bin
sudo note -deploy theme.rocky
rm note.go
rm structs.go
rm config.go
rm commands.go
rm go.sum
rm go.mod
rm theme.rocky
rm build.sh
