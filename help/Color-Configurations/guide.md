# Color Configurations!
___
I'm assuming that you might have stumbled over here by pure curiosity or not, anyways lets get to it!

First off make sure you have Note installed!

You are able to set the text and cursor color to these colors:
- **Red**
- **Green**
- **Yellow**
- **Blue**
- **Purple**
- **Cyan**
- **White**

There are two ways to do this, edit the **colorConfig.json file**, for example say I want to set my text color to Cyan and my cursor color to white I would do this:
```json
{
  "TextColor": [
    {
      "ColorRed":"false",
      "ColorGreen": "false",
      "ColorYellow": "false",
      "ColorBlue": "false",
      "ColorPurple": "false",
      "ColorCyan" : "true",
      "ColorWhite" : "false"
    }
  ],
  "CursorColor": [

    {
      "ColorRed":"false",
      "ColorGreen": "false",
      "ColorYellow": "false",
      "ColorBlue": "false",
      "ColorPurple": "false",
      "ColorCyan" : "false",
      "ColorWhite" : "true"
    }
  ]
}
```
**colorConfig.json** comes with these default values:
```json
{
  "TextColor": [
    {
      "ColorRed":"false",
      "ColorGreen": "false",
      "ColorYellow": "false",
      "ColorBlue": "false",
      "ColorPurple": "false",
      "ColorCyan" : "false",
      "ColorWhite" : "true"
    }
  ],
  "CursorColor": [

    {
      "ColorRed":"false",
      "ColorGreen": "false",
      "ColorYellow": "false",
      "ColorBlue": "false",
      "ColorPurple": "true",
      "ColorCyan" : "false",
      "ColorWhite" : "false"
    }
  ]
}
```
The second option is to use the built in function on note to change the colors,
for example say I wish to set the text color to blue but want to keep my current cursor color I would just run this:
```bash
note Text-Blue Cursor-SAME
```
or if I want to set the Text color to Red and Cursor color to white then I should run this:
```bash
note Text-Red Cursor-White
```
I hope this guide has been useful to anyone who wishes to use it!