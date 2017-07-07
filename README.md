# Grudge

Grudage is a small library to execute code blocks in a secured environment,
Could be used in online judgers. 

I created the library for learning purposes and it's fairly new,
still working on implementing the basic features.

### Installation

```bash
go install github.com/emostafa/grudge
```


### Running Tests

Change directory to the project files

```bash
go test -v
```



### Usage

The usage is failry simple, each supported language has a function with it's
name, and the function starts with an uppercase letter. e.g:

```go
package main


import (
	"fmt"
	"github.com/emostafa/grudge"
)

func main () {
	output, err := grudge.Python("print('Hello Pyton!')", []string{})
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	fmt.Println(output)
}
```


### Supported Languages
 - Python
 - Javascript
 - Ruby
