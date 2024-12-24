# Note - Nightly

To build nightly just compile main.go.

Editor Command Line

ESC key will exit current mode and then enter command line.

To enter Edit Mode type->
:edit or :e

To enter View Mode type->
:view or :v

For help type->
:help or :h

Multiple Modes.

Edit Mode,

Will load current line into the buffer, all edits will be done on the buffer.

Ehh, stack machine for the editor?



POP -ln 5 -col 3
PUSH {RUNE} -ln 3 -col 4


This text is{\n}
being rendered{\n}
line by line.{\n}


View Mode,

Will load text onto screen, limited by size of screen.


I think im going to write a more java centric way of handling the input.

A seperate file for handling its own input. Shouldnt be too hard