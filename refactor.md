I need to actually write legible code for a second.

At this point it is horrible adding parts to this terminal


Best route would be to:

Enable easy text rendering.

Im maybe thinking of a "pages" format.



startup page, similar to NVIM startup page.

Waits for input event then switches page.

Page also handles events. A page can switch to another page, although
upon switching the page the page will cease to exist.


interface Page struct {
    
}