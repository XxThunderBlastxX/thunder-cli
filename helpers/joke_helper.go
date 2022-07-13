package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/XxThunderBlastxX/models"
	"io/ioutil"
	"net/http"
)

//GetRandomJokes helps to print the joke
func GetRandomJokes() {
	url := "https://icanhazdadjoke.com/"
	resByte := getJokeData(url)
	joke := models.Joke{}

	if err := json.Unmarshal(resByte, &joke); err != nil {
		fmt.Printf("could not unmarshal : %v", err)
	}

	fmt.Println("Joke :- " + joke.Joke)
}

//getJokeData helps to fetch joke json from server
func getJokeData(url string) []byte {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Accept", "application/json")

	res, resErr := http.DefaultClient.Do(req)
	if resErr != nil {
		fmt.Println(resErr)
	}

	resByte, resByteErr := ioutil.ReadAll(res.Body)
	if resByteErr != nil {
		fmt.Println(resByteErr)
	}
	return resByte
}
