package app

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/mgutz/ansi"
	"github.com/tomlazar/table"
)

func FormatGenres(genres []struct{ Name string }) string {
	var genreNames []string
	for _, genre := range genres {
		genreNames = append(genreNames, genre.Name)
	}
	return strings.Join(genreNames, ", ")
}

func PrintTable(header []string, rows [][]string, title string) {
	fmt.Printf("### %s ###\n\n", title)
	tab := table.Table{
		Headers: header,
		Rows:    rows,
	}
	writer := os.Stdout
	tab.WriteTable(writer, &table.Config{
		ShowIndex:       true,
		Color:           true,
		AlternateColors: true,
		TitleColorCode:  ansi.ColorCode("white+buf"),
		AltColorCodes: []string{
			"\u001b[44m",
			"\u001b[45m",
		}})

}

func EllipsizeString(str string, maxLen int) string {
	if len(str) > maxLen {
		return str[:maxLen] + "..."
	}
	return str
}

func ConvertString(input string) string {

	input = strings.TrimSpace(input)
	re := regexp.MustCompile(`\s+`)
	singleSpacedString := re.ReplaceAllString(input, " ")
	result := strings.ReplaceAll(singleSpacedString, " ", "+")
	return result
}
