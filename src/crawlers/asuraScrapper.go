package crawlers

import (
	"AMM/src"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func GetLestEpUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		//Faz algo
	}
	defer resp.Body.Close()
	rootNode, err := html.Parse(resp.Body)
	if err != nil {
		//faz algo
	}
	n := src.SearchFirstNodeOccurrence(rootNode, "class", "eph-num", "div")
	log.Println(n.FirstChild.NextSibling.Attr[0].Val) //url of the last epp
}
