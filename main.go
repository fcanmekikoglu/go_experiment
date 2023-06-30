package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/fcanmekikoglu/go_experiment/db"
	"github.com/fcanmekikoglu/go_experiment/types"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Error loading .env file")
		os.Exit(1)
	}

	// Make request
	request, err := http.NewRequest("GET", "https://cat-fact.herokuapp.com/facts", nil)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	request.Header.Set("Accept", "application/json")

	// This is how to add Bearer tokens
	// token := "your-token"
	// request.Header.Set("Authorization", "Bearer "+token)

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

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		fmt.Print("MONGODB_URI not set in .env file")
		os.Exit(1)
	}

	mongoClient, err := db.ConnectMongo(mongoURI)
	if err != nil {
		fmt.Println("Failed to connect MongoDB", err)
		os.Exit(1)
	}

	// Iterate over facts
	for _, fact := range facts {
		err = db.InsertFact(mongoClient, fact, "cat-api", "facts")
		if err != nil {
			fmt.Printf("Failed to insert fact: %s\n", err)
			os.Exit(1)
		}
		// fmt.Printf("Fact %d: %s\n", i+1, fact.Text)
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
