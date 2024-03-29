/*
Copyright 2023 Juan Jose Vargas Fletes

This work is licensed under the Creative Commons Attribution-NonCommercial (CC BY-NC) license.
To view a copy of this license, visit https://creativecommons.org/licenses/by-nc/4.0/

Under the CC BY-NC license, you are free to:

- Share: copy and redistribute the material in any medium or format
- Adapt: remix, transform, and build upon the material

Under the following terms:

  - Attribution: You must give appropriate credit, provide a link to the license, and indicate if changes were made.
    You may do so in any reasonable manner, but not in any way that suggests the licensor endorses you or your use.

- Non-Commercial: You may not use the material for commercial purposes.

You are free to use this work for personal or non-commercial purposes.
If you would like to use this work for commercial purposes, please contact Juan Jose Vargas Fletes at juanvfletes@gmail.com.
*/
package common

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Singleton object for the logger
var logger *Logger = nil

// Color Codes Enum
const (
	COLOR_BLACK int = 30 + iota
	COLOR_RED
	COLOR_GREEN
	COLOR_YELLOW
	COLOR_BLUE
	COLOR_PURPLE
	COLOR_CYAN
	COLOR_WHITE
)

// Style Codes Enum
const (
	STYLE_NO_EFFECT int = iota
	STYLE_BOLD
	STYLE_UNDERLINE
	STYLE_NEG_1
	STYLE_NEG_2
)

// Background Codes Enum
const (
	BG_BLACK int = 40 + iota
	BG_RED
	BG_GREEN
	BG_YELLOW
	BG_BLUE
	BG_PURPLE
	BG_CYAN
	BG_WHITE
)

// Singleton Function for getting the logger
func GetLogger() *Logger {
	if logger == nil {
		logger = &Logger{
			headline: "System",
		}
	}

	return logger
}

// Logs a message
func (l *Logger) Log(section, message string) {
	greenBold := l.getColor(STYLE_BOLD, COLOR_GREEN, BG_BLACK)
	yellowBold := l.getColor(STYLE_BOLD, COLOR_YELLOW, BG_BLACK)
	normalWhite := l.getColor(STYLE_NO_EFFECT, COLOR_WHITE, BG_BLACK)

	fmt.Printf("%s[%s%s%s][%s%s%s]:%s %s\n", greenBold, yellowBold, l.headline, greenBold, yellowBold, section, greenBold, normalWhite, message)
}

// Logs an object message
func (l *Logger) LogObject(section, object any) {
	greenBold := l.getColor(STYLE_BOLD, COLOR_GREEN, BG_BLACK)
	green := l.getColor(STYLE_NO_EFFECT, COLOR_GREEN, BG_BLACK)
	yellowBold := l.getColor(STYLE_BOLD, COLOR_YELLOW, BG_BLACK)
	normalWhite := l.getColor(STYLE_NO_EFFECT, COLOR_WHITE, BG_BLACK)

	b, err := json.MarshalIndent(object, "", "\t")

	if err != nil {
		l.Error("Logger", err.Error())
		return
	}

	fmt.Printf("%s[%s%s%s][%s%s%s]:%s \n%s%s%s\n", greenBold, yellowBold, l.headline, greenBold, yellowBold, section, greenBold, normalWhite, green, b, normalWhite)
}

// Prints a bracketed message
func (l *Logger) BracketLog(message string) {
	greenBold := l.getColor(STYLE_BOLD, COLOR_GREEN, BG_BLACK)
	normalWhite := l.getColor(STYLE_NO_EFFECT, COLOR_WHITE, BG_BLACK)

	fmt.Printf("%s[%s%s%s]%s", greenBold, normalWhite, message, greenBold, normalWhite)
}

// Prints an error
func (l *Logger) Error(section, message string) {
	yellowBold := l.getColor(STYLE_BOLD, COLOR_YELLOW, BG_BLACK)
	redBold := l.getColor(STYLE_BOLD, COLOR_RED, BG_BLACK)
	normalWhite := l.getColor(STYLE_NO_EFFECT, COLOR_WHITE, BG_BLACK)

	fmt.Printf("%s[%s%s%s][%s%s%s]:%s %s\n", yellowBold, redBold, l.headline, yellowBold, redBold, section, yellowBold, normalWhite, message)
}

// Prints a test error
func (l *Logger) TestError(testName string, expected, got any, t *testing.T) {
	greenBold := l.getColor(STYLE_BOLD, COLOR_GREEN, BG_BLACK)
	redBold := l.getColor(STYLE_BOLD, COLOR_RED, BG_BLACK)
	normalWhite := l.getColor(STYLE_NO_EFFECT, COLOR_WHITE, BG_BLACK)
	yellowBold := l.getColor(STYLE_BOLD, COLOR_YELLOW, BG_BLACK)

	t.Errorf("\n[%s%s%s]:\n\t%sGot: %v%s\n\t%sExpected: %v%s\n", yellowBold, testName, normalWhite, redBold, got, normalWhite, greenBold, expected, normalWhite)
}

// Returns a color with a certain format
func (l *Logger) getColor(style, color, bgColor int) string {
	return fmt.Sprintf("\033[%d;%d;%dm", style, color, bgColor)
}
