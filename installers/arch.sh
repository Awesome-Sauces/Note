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
touch themes.rocky
sudo mv -i themes.rocky ~/note

# Grabs Newest version of Note
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/TextEditor.go --output TextEditor.go
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/themes/theme.rocky --output theme.rocky
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/note.go --output note.go
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/structs.go --output structs.go
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/eval.go --output eval.go
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/lexer.go --output lexer.go
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/debug.go --output debug.go
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/commands.go --output commands.go
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/go.mod --output go.mod
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/go.sum --output go.sum
curl https://raw.githubusercontent.com/Awesome-Sauces/Note/main/src/build.sh --output build.sh

sh build.sh
sudo mv -i note /usr/bin
sudo note -deploy theme.rocky
rm note.go
rm structs.go
rm eval.go
rm lexer.go
rm debug.go
rm commands.go
rm go.sum
rm go.mod
rm theme.rocky
rm build.sh
