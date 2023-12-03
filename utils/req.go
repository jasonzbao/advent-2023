package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func MakeRequest(day int) string {
	url := fmt.Sprintf("https://adventofcode.com/2023/day/%v/input", day)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.AddCookie(
		&http.Cookie{
			Name:  "session",
			Value: os.Getenv("AOC_SESSION"),
		},
	)

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func FormatInput(input string) []string {
	return strings.Split(strings.TrimSuffix(input, "\n"), "\n")
}

func FormattedRequest(day int) []string {
	return FormatInput(MakeRequest(day))
}
