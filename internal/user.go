package internal

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"bytes"
	"encoding/json"
)

var LastTabCounter int = 0

var reader = bufio.NewReader(os.Stdin)

/**
* This function is used to read user input from the console indenting the input.
*/
func GetUserInput() string {
	fmt.Print("> " + strings.Repeat("\t", LastTabCounter))
	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("\nError reading input: ", err)
		os.Exit(1)
	}

	return strings.TrimSuffix(text, "\n")
}

/**
* @params str: `the unformatted json content`
* This function formats the given JSON string to be more readable.
*/
func FormatJson(str string) (string, error) {
    var prettyJSON bytes.Buffer
    if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
        return "", err
    }
    return prettyJSON.String(), nil
}

/**
* @params input: `user input`
* This function parses the user input fixing the indentation and GraphQL requirements.
*/
func HandleUserInput(input string) string {
	if strings.Count(input, "}") > 0 {
		if LastTabCounter > 0 {
			LastTabCounter -= 1
		}
	}

	if strings.Count(input, "{") > 0 {
		LastTabCounter += 1
	}

	input = strings.Replace(input, "\"", "\\" + "\"", strings.Count(input, "\""))

	LastTabCounter += strings.Count(input, "\t")

	return strings.Replace(input, "\t", "", strings.Count(input, "\"")) + " "
}
