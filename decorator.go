// Copyright (c) 2017 Pronin S.V.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fmtc

import (
	"strings"
	"unicode/utf8"

	"golang.org/x/net/html"
)

const (
	//CODERESET reset style after this code
	CODERESET = "\033[0m"
)

// An array of tag and ASCI code matches for text
var fontStyle = map[string]string{
	"b":       "1",  // <b> - Bold
	"strong":  "1",  // <strong> - Bold
	"u":       "4",  // <u> - Underline
	"dim":     "2",  // <dim> - Dim
	"reverse": "7",  // <reverse> - Reverse
	"blink":   "5",  // <blink> - Blink
	"black":   "30", // <black> - Black (Dark)
	"red":     "31", // <red> - Red
	"green":   "32", // <green> - Green
	"yellow":  "33", // <yellow> - Yellow
	"blue":    "34", // <blue> - Blue
	"magenta": "35", // <magenta> - Magenta
	"cyan":    "36", // <cyan> - Cyan
	"grey":    "37", // <grey> - Grey (Smokey)
	"white":   "97", // <white> - White
}

// An array of tag and ASCI code matches for background colors
var background = map[string]string{
	"black":   "40",  // Black (Dark)
	"red":     "41",  // Red
	"green":   "42",  // Green
	"yellow":  "43",  // Yellow
	"blue":    "44",  // Blue
	"magenta": "45",  // Magenta
	"cyan":    "46",  // Cyan
	"grey":    "47",  // Grey (Smokey)
	"white":   "107", // White
}

// LIFO stack struct fir HTML tags
type stack struct {
	items []tag
}

// push add element to stack
func (stack *stack) push(value tag) {
	stack.items = append(stack.items, value)
}

// pop delete element from stack
func (stack *stack) pop(tag tag) bool {
	if len(stack.items) > 0 {
		i := len(stack.items) - 1
		if stack.items[i].name == tag.name {
			stack.items = stack.items[:i]
			return true
		}
	}
	return false
}

// tag in this struct we store html-tag name and map of tag attributes
type tag struct {
	name string
	attr map[string]string
}

// decorate string using HTML tags from her
// It returns decorated string
func decorate(str string) string {
	//get io.Reader from string
	reader := strings.NewReader(str)
	//get HTML Tokenizer from io.Reader
	d := html.NewTokenizer(reader)
	//Stack of unclosed HTML tags
	tagsStack := stack{}
	//Result string
	finalString := ""
	for {
		tokenType := d.Next()

		//if str is end, or if error
		if tokenType == html.ErrorToken {

			finalString += CODERESET
			return finalString
		}

		//get current token of str
		token := d.Token()
		switch tokenType {
		//if token type is StartTag (like <tagname>)
		case html.StartTagToken:
			oneTag := tag{name: token.Data}
			//try to get tag attributes
			if len(token.Attr) > 0 {
				tagAttr := make(map[string]string)
				for _, attr := range token.Attr {
					tagAttr[attr.Key] = attr.Val
				}
				oneTag.attr = tagAttr
			}
			//try to get tag ANSI code
			tagCode := getASCICode(oneTag)
			if utf8.RuneCountInString(tagCode) > 0 {
				//if ANSI code exists, then we add them to string
				finalString += tagCode
				//and add current tag in stack of opened tags
				tagsStack.push(oneTag)
			} else {
				//if tag ANSI code not exists, then add to final string tag
				finalString += token.String()
			}
		//if curent token is text
		case html.TextToken:
			//then add this text to final string
			finalString += token.Data
		//if current token is SelfClosingTag (like <br />)
		case html.SelfClosingTagToken:
			//add tag to string
			finalString += token.String()
		//if current token is EndTag (like </tagname>)
		case html.EndTagToken:
			oneTag := tag{name: token.Data}
			//try to pop last tag from stack (Success if the last stack tag matches the current closing tag)
			if tagsStack.pop(oneTag) == true {
				finalString += CODERESET
				finalString += applyOpenedTags(tagsStack)
			} else {
				finalString += token.String()
			}
		}
	}
}

// decorateIface decorator for an empty interface
func decorateIface(a *[]interface{}) {
	for i, x := range *a {
		if s, ok := x.(string); ok {
			(*a)[i] = decorate(s)
		}
	}
}

// applyOpenedTags apply to string not closed tagsArr
// It returns string with ASCI codes of not closet tags
func applyOpenedTags(tagsStack stack) string {
	if len(tagsStack.items) == 0 {
		return ""
	}
	tagsStr := ""
	for _, tag := range tagsStack.items {
		tagCode := getASCICode(tag)
		if utf8.RuneCountInString(tagCode) > 0 {
			tagsStr += tagCode
		}
	}
	return tagsStr
}

// getASCICode get ASCI code of font style or background color
// It returns the code of current tag
func getASCICode(tag tag) string {
	code := ""
	if tag.name == "bg" {
		if color, ok := tag.attr["color"]; ok {
			code = background[color]
		}
	} else {
		code = fontStyle[tag.name]
	}
	if utf8.RuneCountInString(code) > 0 {
		code = "\033[" + code + "m"
	}
	return code
}
