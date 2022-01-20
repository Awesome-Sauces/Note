# Setup
___
## **NOTE: GoEdit is not compatible with Windows**
## **Installing Golang**
[Mac/Linux](https://go.dev/dl/) 
(**If on Arch Linux just run this command** ```sudo pacman -S go```)


## **Steps**
Download any version you like of GoEdit!
Compile **GoEdit.go** with ```go build GoEdit.go```
    * Open the directory in Terminal, which you built **GoEdit.go** in.
Run the following command in the directory which you compiled GoEdit.go.
```shell
sudo cp GoEdit /usr/bin
```
There you're set and done! To use GoEdit just move into any directory you wish and run it with the following format
```shell
GoEdit -a FILENAME
```
To append/edit the given file
or
``` shell
GoEdit -dc FILENAME
```
To delete all file contents!
Now you have successfully installed the GoEdit version of your choice!