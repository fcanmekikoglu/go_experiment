package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type User struct {
	ID        string `json:"_id"`
	Version   string `json:"_v"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Email     string `json:"email"`
	IsAdmin   bool   `json:"isAdmin"`
	Name      struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"name"`
	Google struct {
		ID           string `json:"id"`
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	} `json:"google"`
}

type Fact struct {
	ID        string `json:"_id"`
	Version   int    `json:"_v"`
	User      string `json:"user"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Deleted   bool   `json:"deleted"`
	Source    string `json:"source"`
	Type      string `json:"type"`
	Text      string `json:"text"`
	Status    struct {
		Verified  bool   `json:"verified"`
		SentCount int    `json:"sentCount"`
		Feedback  string `json:"feedback"`
	} `json:"status"`
}

func main() {
	// Make request
	response, err := http.Get("https://cat-fact.herokuapp.com/facts")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

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
	var facts []Fact
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
