package main

import (
	"encoding/json"
	"flag"
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

	// Initialize mongo client
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

	// Handling flags
	dumpFlag := flag.Bool("dump", false, "Dumps target collection")
	dbFlag := flag.String("db", "none", "target database")
	collectionFlag := flag.String("collection", "none", "target collection")

	flag.Parse()
	if *dumpFlag == false {

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

		// Iterate over facts
		for _, fact := range facts {
			err = db.InsertFact(mongoClient, fact, "cat-api", "facts")
			if err != nil {
				fmt.Printf("Failed to insert fact: %s\n", err)
				os.Exit(1)
			}
			// fmt.Printf("Fact %d: %s\n", i+1, fact.Text)
		}

	} else if *dumpFlag && (*dbFlag == "none" || *collectionFlag == "none") {
		fmt.Println("You should provide database and collection name if you want to use --dump")
		flag.PrintDefaults()
		os.Exit(1)
	} else if *dumpFlag {
		// call mongo dump
		err = db.DumpCollection(mongoClient, *dbFlag, *collectionFlag)
		if err != nil {
			fmt.Printf("Error while dumping db %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Successfully dumped collection %s%s\n", *dbFlag, *collectionFlag)
		os.Exit(0)
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
