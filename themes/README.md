# Note color themes

Note color themes are used to give a little more lively hood to your note text editor experience.
Although these themes won't provide vs-code extension like abilities it will allow you to be able to make
dense color themes to further customize your experience and even provide some syntax highlighting.

## Standard Functions
```python
# Prints text pre load into note instance
printf(text)
# Sets the linenumber color and background color
LineColors(lineNumberColorHex, styleCanBeEmpty, BackgroundColorHex)
# Sets background color
Background(colorHex)
# Adds a new keyword
NewKeyword(extension, word, colorHex)
# Adds a time delay in seconds
sleep(5)
```
## Examples
This will make a variable and print each iteration
```javascript
// Initialize variable with $ symbol
int $num = 5;

// Call variable with $ symbol
loop($num){
    // Print loop iteration
    printf($loop);
}

```
This will use the Background, LineColors and sleep functions
```javascript
string $bgColor = "#303030";

string $numColor = "#b2b2b2";
string $lineColorbg = "#4e4e4e";
string $style = "b";

Background($bgColor);

LineColors($numColor, $style, $lineColorbg);

sleep(1);

printf("Number Color ", $numColor," Style ", $style, " Line Color bg ", $lineColorbg);
printf("BackGround Color " + $bgColor);
```
This will be a decent color theme example
```javascript
string $y = "#303030";

string $numColor = "#b2b2b2";
string $shadow = "#4e4e4e";
string $style = "b";

int $loc = 5;

Background($y);

LineColors($numColor, $style, $shadow);

printf("Background set to " $y);


list $keywords = ["for", "while", "def", "class", "print", "fmt", "Println", "'"];
list $colors = ["#ff87ff", "#ff87ff", "#ff87ff", "#ff87ff", "#5f87d7", "#5f87d7", "#ff87ff", "#5f87d7"];

loop($keywords){
	NewKeyword(".py", $keywords[$loop], $colors[$loop])	
}
```
## Deploying and Testing
To test a script please run this command
```
note -script filename
```
To deploy a color theme to note please do this command
```
note -deploy filename
```
It will automatically set it as the current running theme.

## Adding functionability
For now rocky is under development but will recieve much more additions in the future but for now to add
more to the rocky language if you wish please edit the eval.go or lexer.go and make a pull request. If you
happen to find any bugs with rocky please make an issue and I will try my best to solve it!

If you have any recommendations or bugs about Note, please feel free to open an issue and report them.

![Note-logos_black-Logo](https://user-images.githubusercontent.com/78565561/150656857-c89e1528-9f4b-4df2-bd51-c43456c720c0.png)
