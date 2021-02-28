# Water Jug Riddle

This application solves the water jug riddle arithmetically, determining first the optimal solution and subsequently streaming each state change via channels.
Consuming from the stream and acknowledging the display of each state. I choose using channels to stream the states efficiently avoiding extra memory storage.
The display of each state is rendered in the terminal but could be potentially extended to any other streaming display or format.

## Make
To run the tests and build the executable in a single step using make.

```bash
make
```

## Run
To simply run the program without building the binary

```bash
go run cmd/main.go
```
