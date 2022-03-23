package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

//IamDocsUrl is the base reference for all aws service documentation
const IamDocsUrl = "https://docs.aws.amazon.com/service-authorization/latest/reference/reference_policies_actions-resources-contextkeys.html"

//IamDocsBaseUrl is used for rendering FQDNs for all aws services as relative Urls are used within the site
const IamDocsBaseUrl = "https://docs.aws.amazon.com/service-authorization/latest/reference/"

func main() {
	c := colly.NewCollector()

	c.OnHTML("div[class=highlights]", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, el *colly.HTMLElement) {
			fmt.Printf("%+v %+v \n", el.Text, fixIAMDocsUrls(el.Attr("href")))
		})
	})

	c.Visit(IamDocsUrl)
}

func fixIAMDocsUrls(href string) string {
	joinedUrl := strings.Replace(href, "./", "", 1)
	return fmt.Sprintf("%s%s", IamDocsBaseUrl, joinedUrl)
}
