package common

import (
	"encoding/json"
	"fmt"
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

// Returns a color with a certain format
func (l *Logger) getColor(style, color, bgColor int) string {
	return fmt.Sprintf("\033[%d;%d;%dm", style, color, bgColor)
}
