// Package colorizer provides utilities for applying ANSI colors to text output in terminals.

package colorizer

import (
    "fmt"
)

// Color represents different ANSI colors
type Color int

// Constants for different ANSI colors
const (
    Black       Color = iota // Black color
    RedColor                // Red color
    GreenColor              // Green color
    YellowColor             // Yellow color
    BlueColor               // Blue color
    MagentaColor            // Magenta color
    CyanColor               // Cyan color
    WhiteColor              // White color
)

// Colorize applies ANSI color codes to the input text based on the specified color and intensity.
// The intensity parameter ranges from 1 to 3, with 1 being the least intense
// and 3 being the most intense.
func Colorize(text string, color Color, intensity int) string {
    code := colorCode(color, intensity)
    return fmt.Sprintf("\033[%dm%s\033[0m", code, text)
}

// colorCode returns ANSI escape code for given color and intensity.
func colorCode(color Color, intensity int) int {
    base := 30 // Base ANSI code for foreground colors
    if intensity < 1 {
        intensity = 1
    } else if intensity > 3 {
        intensity = 3
    }

    // Calculate code for the specified color and intensity
    code := base + int(color)*3 + intensity - 1
    return code
}

// Red applies the red ANSI color to the input string.
func Red(text string, intensity int) string {
    return Colorize(text, RedColor, intensity)
}

// Blue applies the blue ANSI color to the input string.
func Blue(text string, intensity int) string {
    return Colorize(text, BlueColor, intensity)
}

// Green applies the green ANSI color to the input string.
func Green(text string, intensity int) string {
    return Colorize(text, GreenColor, intensity)
}

// Yellow applies the yellow ANSI color to the input string.
func Yellow(text string, intensity int) string {
    return Colorize(text, YellowColor, intensity)
}

// Cyan applies the cyan ANSI color to the input string.
func Cyan(text string, intensity int) string {
    return Colorize(text, CyanColor, intensity)
}

// CustomColor applies a custom ANSI color, specified by colorNumber, to the input string.
// The colorNumber parameter should be an integer representing a custom ANSI color.
func CustomColor(text string, colorNumber int, intensity int) string {
    return Colorize(text, Color(colorNumber), intensity)
}
