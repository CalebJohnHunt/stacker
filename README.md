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

## Basic Usage

```go
/* main.go */
import "github.com/CalebJohnHunt/stacker"

type myModel struct{}

func (m myModel) Init() tea.Cmd {
  // This func will be called when the model is placed on top of the stack
  return nil
}

func (m myModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "enter": // push new scene
      return m, stacker.AddNewScene(myModel{})
    case "esc": // pop current scene
      return m, stacker.PopScene()
    }
  case myModel: // A previously pushed scene popped itself and now we're top-of-stack!
    // ... do something with myModel
    return m, nil
  }
  return m, nil
}

func (m myModel) View() string {
  // This model's view will be used when it is on top of the stack
  return "View of myModel"
}

func main() {
  // make stacker with your model as the lowest layer
  tea.NewProgram(stacker.NewStacker(myModel{})).Run()
}
```

## API

| API                           | Explanation |
|:------------------------------|:--|
| `NewStacker(tea.Model)`       | Create a new stacker with your `tea.Model` as the bottom layer. |
| `AddScene(tea.Model) tea.Cmd` | Push a new scene (model) onto the stack. This means the current model won't be shown and won't receive updates (until the pushed scene pops itself) |
| `PopScene() tea.Cmd`          | Pop the current scene and provide it to the new top-of-stack scene. This is done by way of `tea.Msg`. See the [basic usage](#basic-usage) for an example. |
| `PopSceneSilent() tea.Cmd`    | Pop the current scene but do not provide any indication to the new top-of-stack scene. |

## Todo

+ [X] Return the model when the scene is popped
+ [X] Push a scene which, when popped, will not alert the pusher; or perhaps popping can be silent... That sounds easier
+ [ ] Put stackers on stackers
