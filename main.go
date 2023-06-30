package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/fcanmekikoglu/go_experiment/types"
)

func main() {
	// Make request
	request, err := http.NewRequest("GET", "https://cat-fact.herokuapp.com/facts", nil)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	request.Header.Set("Accept", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)

	if response.StatusCode != 200 {
		fmt.Println("Error!, Status is not 200!")
		os.Exit(1)
	}

	// Get response body
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// Unmarshal facts and extract
	var facts []types.Fact
	json.Unmarshal(responseData, &facts)

	// Iterate over facts
	for i, fact := range facts {
		fmt.Printf("Fact %d: %s\n", i+1, fact.Text)
	}

	// // Create file
	// file, err := os.Create("response.json")
	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	os.Exit(1)
	// }

	// // Write response to file & close
	// _, err = file.Write(responseData)
	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	os.Exit(1)
	// }
	// file.Close()

	// // Print if success
	// fmt.Println("JSON data saved to response.json")
}
