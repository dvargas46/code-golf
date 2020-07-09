package routes

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/code-golf/code-golf/hole"
)

// Sitemap serves GET /sitemap.xml
func Sitemap(w http.ResponseWriter, r *http.Request) {
	type URL struct {
		Loc string `xml:"loc"`
	}

	type URLSet struct {
		XMLName xml.Name `xml:"http://www.sitemaps.org/schemas/sitemap/0.9 urlset"`
		URLs    []URL    `xml:"url"`
	}

	sitemap := URLSet{
		URLs: []URL{
			{"https://code.golf"},
			{"https://code.golf/about"},
			{"https://code.golf/ideas"},
			{"https://code.golf/stats"},
		},
	}

	for _, hole := range hole.List {
		sitemap.URLs = append(
			sitemap.URLs, URL{"https://code.golf/" + url.PathEscape(hole.ID)})
	}

	w.Header().Set("Content-Type", "text/xml; charset=utf-8")
	w.Write([]byte(xml.Header))

	if err := xml.NewEncoder(w).Encode(sitemap); err != nil {
		panic(err)
	}
}
