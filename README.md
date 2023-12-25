## find couple game

Is engine of game for find a pair of items.

Game is storing more data as secret. 
You can't get item content without click to that at least once.

Game working with any correct sizes and items. For universal game stores only `item_id`.
Render of cards is client side task.

Rules:
>`S = BoardHeight*BoardWidth ` \
>**must**: `S % 2 == 0` \
>**must**: `len(items) == S/2` 

Items has 3 states:
1. Hided - item content is hided, actually.
2. Opened - active item (last click).
3. Guessed - user already find pair for this item. 

## Execute game
Base game object stores in `domain/game`.

You need create game object via `game.New` and call `Init` method.
If you don't initialize game you will get panic.

For process click call method `MakeClick`.
Game stores last click for state, when one cards is opened.

If you click on not hided card it will do nothing.

When you find last pair method return as result status `game.Finish`
it means that game is end. For execute with similar params you can call `Init` method.


## Output
In pkg/visual directory stores some interfaces for rendering.

At the now it not universal interfaces and bad for many outputs. 
Improving this later.


In desk/console stores full done game process. It also have not good architecture.


## Problems
1. Rendering wrong step.
2. Really universal mechanism for rendering
