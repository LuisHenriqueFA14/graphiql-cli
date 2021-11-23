package main

import (
	"fmt"
	"os"

	"github.com/LuisHenriqueFA14/graphiql-cli/internal"
)

var fullQuery string
var url string

func main() {
	fmt.Println("What is the GraphQL URL ?")
	url = internal.GetUserInput()

	fmt.Println("Write your query: (stop with ';')")

	for {
		query := internal.GetUserInput()

		if query == ";" {
			break
		}

		fullQuery += internal.HandleUserInput(query)
	}

	fmt.Println("\nRunning query...")

	result, err := internal.FormatJson(internal.RunQuery(url, fullQuery))

	if err != nil {
		fmt.Println("\nError formatting JSON: ", err)
		os.Exit(1)
	}

	fmt.Println("Response: \n" + result)
}
