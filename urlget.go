package main

import "sync"

//Links - structure to hold my links in. with mutex to allow safe writes
type Links struct {
	entries map[string]string
	mux     *sync.Mutex
}

// NewLinkmap Constructs a new NewLinkMap.
func NewLinks() *Links {
	return &Links{
		entries: make(map[string]string),
		mux:     &sync.Mutex{},
	}
}

// GetLinks crawls a start URL for all links and assets and builds
// a links struct with pages and assets per crawled link.
func GetLinks(url string) (*Links, error) {
	links := NewLinks()
	return links, nil
}

// need a function to return a string to return to client
// obviously a JSON string...

func (s *Links) PrintLinks() string {
	//all External
	strRet := "All External Links\n"
	for key, value := range s.entries {
		if value == "External" {
			strRet += "\n<p>"
			strRet += key
			//fmt.Printf("%s\n", key)
		}
	}
	strRet += ""
	//all assets
	strRet += "All Assets\n"
	for key, value := range s.entries {
		if value == "Asset" {
			strRet += "\n<p>"
			strRet += key
		}
	}
	strRet += ""
	//all pages
	strRet += "All Pages\n"
	for key, value := range s.entries {
		if value == "Page" {
			strRet += "\n<p>"
			strRet += key
		}
	}
	return strRet
}
