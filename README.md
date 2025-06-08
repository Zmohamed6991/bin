# Numbers to Words

A simple Go program that converts numbers (0 to 100,000) into their English words representation.

---

## Features

- Converts integers from 0 up to 100,000 into English words.
- Handles hundreds and thousands with proper formatting.
- Includes a command-line interface (CLI) for easy usage.
- Includes unit tests using table-driven testing.

---

## Installation

Make sure you have [Go installed](https://golang.org/dl/).

Clone this repository:
git clone https://github.com/zmohamed6991/bin/numbers-to-words.git
cd numbers-to-words

Initialise the Go module (if not already initialised):

go mod init github.com/zmohamed6991/bin/numbers-to-words
Download dependencies:
go mod tidy

Run in the terminal to build the executables:
go build -o numbers-to-words

Run the program from the command linevfollowed by the number you want to convert:

./bin/numbers-to-words  65
Output:
sixty-five

Testing
Run all tests with:
go test -v


