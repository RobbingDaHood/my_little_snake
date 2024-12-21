# my_little_snake

A small game written in golang.

Use wasd to move the snake (you have to hit enter in between each move).

See "configurations.go" for configuration options.

# Concurrency

The game state is owned by the function in "state.go" and ensures:

1. That only one event changes the game state at a time.
1. Because the handleSnake only writes to data that updating direction does not, then they are not overwriting each
   other.
2. There is a small risk that multiple change direction events are sent at the same time only persisting the last
   received. This could be fixed with a mutex but is likely not worth the performance hit because that is based on human
   input.
3. The "render" gets a copy of the game state and renders that. This is done to ensure that the game state is not
   changed
   while rendering and still being open for changes to the state. (If the game state were much larger, then there would
   likely be a performance hit on copying it all so often)

# Linting 

Did perform:
```shell
go vet ./...
staticcheck ./... # Before I updated go to newest version. It still depends on go1.22.2
```