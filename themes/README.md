# Note color themes

Note color themes

## Simple Commands
```
Usage: note [-arg]
    note [--help]
    note [filename]
    note [-txt] [color]
    note [-bg] [color]
Color list:
    [white]
    [red]
    [black]
    [blue]
    [green]
    [purple]
    [yellow]
```
## Example
This command opens 2 files
```
note main.py helper.py
```
This command sets the background color to black
```
note -bg black
```
This command sets the text color to white
```
note -txt white
```
## Installing
To install note there are various ways to do so, one way that is simple is downloading the proper installer script **(Installer scripts are located in installers directory and end with the .sh extension, for a deep guide on installing [click here](https://github.com/Awesome-Sauces/Note/tree/main/installers#readme))** for your operating system or linux distrobution. I still have to expand the installer scripts a bit and test on more distro's. Compiling from source is another one, there is even a bash script to help you do so! Located in the src directory all you have to do is run this command to build the source **(Assuming you have golang already installed)**
```
sh build.sh
```
If you have any recommendations or bugs about Note, please feel free to open an issue and report them.

A text editor for the Linux Terminal! (Mainly compatible with Arch, because I made it on there)


![Note-logos_black-Logo](https://user-images.githubusercontent.com/78565561/150656857-c89e1528-9f4b-4df2-bd51-c43456c720c0.png)
