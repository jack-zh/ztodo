gotodo

A simple command-line todo list written in Go.

Installation:

1. Install Go:
	http://golang.org/doc/install.html

2. Install todo:
	$ go install github.com/jack-zh/gotodo

Usage:

1. Add a task to the list:


	$ todo one
	$ todo two


2. List first tasks:


	$ todo
	 one


3. List outstanding tasks:


	$ todo ls
      1: two
      0: one


4. Remove the first task from the list:


	$ todo rm


5. Remove a first task from the list:


	$ todo rm 0


Next version:

1. add command `todo add`
2. add command `todo done`
3. add command `todo do`
3. show list like this:
    
    2   [Future]:  three 
    1   [Doing]:   two
    0   [Done]:    one
    


