// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"

// 	algorithmia "github.com/algorithmiaio/algorithmia-go"
// )

// func main() {
// 	var apiKey = "sim5vLefek/igv1k9D05HTes1f51"

// 	// Create the Algorithmia client object
// 	var client = algorithmia.NewClient(apiKey, "")
// 	algo, _ := client.Algo("web/SiteMap/0.1.7")
// 	resp, _ := algo.Pipe("http://myfave.com")
// 	response := resp.(*algorithmia.AlgoResponse)
// 	res, _ := json.Marshal(response)
// 	ioutil.WriteFile("linksgo.json", res, 0644)
// 	fmt.Println(response.Result)

// 	// links := []string{"https://myfave.com/surabaya/eat?category_ids=2", "https://myfave.com/surabaya/eat?category_ids=20"}

// 	// for _, l := range links {
// 	// 	analgo, _ := client.Algo("web/AnalyzeURL/0.2.14")
// 	// 	anresp, _ := analgo.Pipe(l)
// 	// 	anresponse := anresp.(*algorithmia.AlgoResponse)
// 	// 	fmt.Println(anresponse.Result)
// 	ioutil.WriteFile("linksgo2.json", res, 0644)

// }