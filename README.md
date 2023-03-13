# Stacker

Scene stacking for the [BubbleTea](https://github.com/charmbracelet/bubbletea).

## I already know BubbleTea. What is this?

A way to stack scenes on top of one another and interact with only the top one. For example:

Here is a "main menu" scene:

```txt
Choose a game:
 > Snake
   TicTacToe
   Tetris
````

And imagine there is a scene for each game listed.  
When a user selects a game, they would expect to play the game. Then when they're done,
they should be brought back to the main menu. For the sake of example, imagine the Snake game also
has an "Options" menu which users would again expect to be able to enter and exit (back into their
Snake game). This could be conceptually thought of as a stack of scenes!

```txt
The "stack" grows rightwards

- Program starts
| Main menu |
- User selects Snake
| Main menu | Snake |
- User plays Snake for a while...
| Main menu | Snake |
- User check options for Snake
| Main menu | Snake | Options |
- User saves their options (and returns to Snake)
| Main menu | Snake |
- User finishes playing snake
| Main menu |
- User selects TicTacToe
| Main menu | TicTacToe |
... and so on ...
```

Stacker is a framework to place "scenes" on top of each other. The "stacker" manages the stack
of scenes. Only the top scene is shown and receives updates. The top-most scene can "push" a new scene
on top or "pop" itself off of the stack.

## API

**TODO**  
This API is not even close to functioning yet. I think it can be used, but I also wouldn't be surprised
if I've neglected to export a useful/necessary function.

## Todo

+ [X] Return the model when the scene is popped
+ [ ] Push a scene which, when popped, will not alert the pusher; or perhaps popping can be silent... That sounds easier
+ [ ] Put stackers on stackers
