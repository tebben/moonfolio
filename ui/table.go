package ui

import (
	"fmt"
	"unicode/utf8"
)

type columnDefinition struct {
	Header        string
	Size          int
	Enabled       bool
	StylingHeader func() []string
	StylingBody   func(interface{}) []string
}

type columnText struct {
	Text    string
	Styling []string
	Length  int
}

func getRowHeaderStyling() []string {
	return []string{ColorWhite, BoldStart}
}

func getRowColumnDefaultStyling(input interface{}) []string {
	return []string{ColorGray, BoldStart}
}

func getRowChangeStyling(input interface{}) []string {
	switch i := input.(type) {
	case float64:
		return []string{getChangeColorStyle(i), BoldStart}
	default:
		return []string{}
	}
}

func createColumnString(columns []columnText) string {
	totalString := ""
	for _, column := range columns {
		stringToAppend := column.Text
		stringLength := utf8.RuneCountInString(stringToAppend)

		if stringLength > column.Length {
			stringToAppend = stringToAppend[:column.Length]
		} else if stringLength < column.Length {
			spacesToAdd := column.Length - stringLength
			for i := 0; i < spacesToAdd; i++ {
				stringToAppend = fmt.Sprintf("%s ", stringToAppend)
			}
		}

		for i := len(column.Styling) - 1; i >= 0; i-- {
			stringToAppend = fmt.Sprintf("%s%s", column.Styling[i], stringToAppend)
		}

		stringToAppend = fmt.Sprintf("%s%s", stringToAppend, Reset)
		totalString = fmt.Sprintf("%s %s", totalString, stringToAppend)
	}

	return totalString
}

func getColumnSpacer(length int) string {
	return createSpacer(length, ColorGray)
}

func getColumnHeadSpacer(length int) string {
	return createSpacer(length, ColorCyan)
}

func createSpacer(length int, color string) string {
	spacer := fmt.Sprintf("%s–", color)
	for i := 0; i < length; i++ {
		spacer = fmt.Sprintf("%s–", spacer)
	}
	return fmt.Sprintf("%s%s", spacer, Reset)
}
