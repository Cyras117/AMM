package src

import (
	"strings"

	"golang.org/x/net/html"
)

func isIn(phrase string, str string) bool {

	p := strings.ToLower(phrase)
	s := strings.ToLower(str)

	lPhr := len(phrase)
	lStr := len(str)

	if lStr < lPhr {
		return false
	}

	for i := 0; i < lStr; i++ {
		if p[0] == s[i] {
			for j := i; j < lStr; j++ {
				if s[j] == p[j-i] {
					if j-i == lPhr-1 {
						return true
					}
				} else {
					break
				}
			}
		}
	}
	return false
}

func checkMatch(node *html.Node, atrr, value, data string) bool {
	if node.Attr == nil {
		return false
	}
	if !strings.EqualFold(strings.ToLower(node.Data), strings.ToLower(data)) {
		return false
	}
	for i := 0; i < len(node.Attr); i++ {
		natrr := strings.ToLower(node.Attr[i].Key)
		nvalue := strings.ToLower(node.Attr[i].Val)
		atr := strings.ToLower(atrr)
		val := strings.ToLower(value)
		if natrr == atr {
			if isIn(val, nvalue) {
				return true
			}
		}

	}
	return false
}

func SearchNodeByData(node *html.Node, data string) *html.Node {
	nodeResult := node

	if nodeResult.Data == data {
		return nodeResult
	}

	if node.FirstChild != nil {
		nodeResult = SearchNodeByData(node.FirstChild, data)
		if nodeResult != nil {
			return nodeResult
		}
	}

	if node.NextSibling != nil {
		nodeResult = SearchNodeByData(node.NextSibling, data)
		if nodeResult != nil {
			return nodeResult
		}
	}

	return nil
}

func SearchNodeByAtrr(node *html.Node, atrr string, value string, result *[]html.Node, data string) {
	if checkMatch(node, atrr, value, data) {
		*result = append(*result, *node)
	}

	if node.FirstChild != nil {
		SearchNodeByAtrr(node.FirstChild, atrr, value, result, data)
	}

	if node.NextSibling != nil {
		SearchNodeByAtrr(node.NextSibling, atrr, value, result, data)
	}
}

func SearchFirstNodeOccurrence(node *html.Node, atrr string, value string, data string) *html.Node {
	var result []html.Node
	SearchNodeByAtrr(node, atrr, value, &result, data)
	if result == nil {
		return nil
	}
	return &result[0]
}
