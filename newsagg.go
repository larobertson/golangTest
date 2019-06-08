package main

import ("fmt"
				"net/http"
				"io/ioutil"
				"encoding/xml")

// type SitemapIndex struct {
// 	Locations []Location `xml:"sitemap"`
// 	// That [] is a slice of type Location
// 	// xml:"sitemap" references a <sitemap> xml tag
// 	// This essentially tells us to extract the value from that tag
// }

// type Location struct {
// 	Loc string `xml:"loc"`
// 	// same here as above, we want to parse a <loc> tag
// }

// func (l Location) String() string {
// 	return fmt.Sprintf(l.Loc)
// }

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
	// Refactoring SitemapIndex struct
	// tells it to look at sitemap then at loc
	// tells us that we want a slice of type string
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	// fmt.Println(s.Locations)
	for _, Location := range s.Locations {
		fmt.Printf("\n%s", Location)
	}
}