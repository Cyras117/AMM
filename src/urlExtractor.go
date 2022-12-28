package src

import (
	"log"

	"golang.org/x/net/html"
)

func GetAllTheMangasUrl() []string {
	const configScope = "mangaFolder"
	const configAttrKey = "attributeKey"
	const configAttrValue = "attributeValue"
	const configTagName = "tagName"

	var readingNodes []html.Node
	var res []string

	rootNode := Read_HtmlFile()
	searchConfig := loadConfig()

	if rootNode == nil {
		log.Println("rootNode is nil")
		return nil
	}

	mangaFolderNode := SearchFirstNodeOccurrence(rootNode, searchConfig[configScope][configAttrKey],
		searchConfig[configScope][configAttrValue], searchConfig[configScope][configTagName])

	SearchNodeByAtrr(mangaFolderNode, "href", "https://", &readingNodes, "a")

	for _, node := range readingNodes {
		res = append(res, node.Attr[0].Val)
	}

	return res
}
