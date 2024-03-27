# Game of Life

Conway's Game of Life implementation in Go using the [Ebiten engine](https://github.com/hajimehoshi/ebiten)

# Installing

```bash
go mod download
```

# Running

```bash
go run main.go
```

There is also a way to generate cells randomly:

```bash
go run main.go --random <chance>
```

Where `<chance>` is an integer number between 0 and 100, e.g.

```bash
go run main.go --random 30
```

# Controls

- `P`: Pause/Resume
- `I`: Interactive Mode (left mouse click to create cells)
- `R`: Clear cells
- `A`: Generate cells randomly
- `+`: Increase zoom
- `-`: Decrease zoom
