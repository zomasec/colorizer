[![Go Reference](https://pkg.go.dev/badge/github.com/zomasec/colorizer.svg)](https://pkg.go.dev/github.com/zomasec/colorizer) [![Go Report Card](https://goreportcard.com/badge/github.com/zomasec/colorizer)](https://goreportcard.com/report/github.com/zomasec/colorizer)

# Colorizer v1.0.0

Colorizer is a Go package that provides utilities for colorizing text output in terminals using ANSI colors, Developed by Hazem El-Sayed (ZomaSec).

## Installation

To use Colorizer in your Go project, you can simply run:

```bash
go get -u github.com/zomasec/colorizer
```
## Usage

Colorizer offers functions to easily colorize text output in your terminal applications. Here's how you can use it:

```go
package main

import (
    "github.com/zomasec/colorizer"
)

func main() {
    // Colorize text with red color
    redText := colorizer.Red("Hello, world!", 1)
    // Print the colored text
    println(redText)
}

```

## Available Colors
Colorizer supports the following ANSI colors:

- Black
- Red
- Green
- Yellow
- Blue
- Magenta
- Cyan
- White

## Custom Colors
You can also use custom colors by specifying a color number between 0 and 255.
```go
Copy code
package main

import (
    "github.com/zomasec/colorizer"
)

func main() {
    // Colorize text with custom color (color number 200)
    customText := colorizer.CustomColor("Hello, world!", 200, 1)
    // Print the colored text
    println(customText)
}
```

## Intensity
You can adjust the intensity of the colors from 1 to 3.
```go
package main

import (
    "github.com/zomasec/colorizer"
)

func main() {
    // Colorize text with red color and intensity 2
    intenseRedText := colorizer.Red("Hello, world!", 2)
    // Print the colored text
    println(intenseRedText)
}
```

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author
- @ZomaSec





