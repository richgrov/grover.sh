package main

import "strings"

const AnsiReset = "\033[0m"
const AnsiRed = "\033[0;31m"
const AnsiBlue = "\033[0;36m"
const AnsiGray = "\033[0;37m"

func calcPaddingToCenter(lineWidth int, viewportWidth int) string {
	return strings.Repeat(" ", viewportWidth/2-lineWidth/2)
}

func centerText(text string, viewportWidth int) string {
	textLen := len([]rune(text))
	return calcPaddingToCenter(textLen, viewportWidth) + text
}
