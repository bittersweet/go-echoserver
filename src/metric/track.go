package metric

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

const hostname = "https://api.metric.io/"
const api_key = "864199c6ddac7fd0e2434021d56fc074"

func Track(metric string) {
	response, err := http.Get(hostname + "track?api_key=" + api_key + "&metric=" + metric)

	if err != nil {
		fmt.Printf("%s", err)
	} else {
		contents, _ := ioutil.ReadAll(response.Body)

		fmt.Println("TRACK:", string(contents))
	}
}
