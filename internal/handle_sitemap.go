package internal

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
)

func (s *Server) handleSitemap() http.HandlerFunc {
	var sitemap []byte

	type SitemapQuery struct {
		Uri        string `xml:"loc"`
		Changefreq string `xml:"changefreq,omitempty"`
		Priority   string `xml:"priority,omitempty"`
	}

	type SitemapPage struct {
		XMLName xml.Name       `xml:"http://www.sitemaps.org/schemas/sitemap/0.9 urlset"`
		Data    []SitemapQuery `xml:"url"`
	}

	// First run, generate it!
	if len(sitemap) == 0 {
		fuentes, err := s.fr.List()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error listando fuentes: %s\n", err.Error())
			os.Exit(1)
		}

		var parroquias = make(map[string]int)
		for _, fuente := range fuentes {
			if p, ok := parroquias[fuente.Parroquia]; ok {
				parroquias[fuente.Parroquia] = p + 1
			} else {
				parroquias[fuente.Parroquia] = 1
			}
		}

		// Páginas de fuentes
		var pages = []SitemapQuery{}
		for _, fuente := range fuentes {
			pages = append(pages, SitemapQuery{"/fuentes/" + fuente.Id, "yearly", "0.3"})
		}

		for parroquia := range parroquias {
			pages = append(pages, SitemapQuery{"/parroquias/" + parroquia, "yearly", "0.5"})
		}

		pages = append(pages, SitemapQuery{"/", "monthly", "0.9"})

		var domain = os.Getenv("DOMAIN")
		for i, p := range pages {
			p.Uri = domain + p.Uri
			pages[i] = p
		}

		sitemap, err = xml.MarshalIndent(SitemapPage{Data: pages}, "", "\t")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error generando sitemap: %s\n", err.Error())
			os.Exit(1)
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/xml")
		w.Write([]byte(xml.Header))
		w.Write(sitemap)
	}
}
