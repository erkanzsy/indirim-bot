package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"io"
	"log"
	"net/http"
)

const BaseUrl = "https://www.hepsiburada.com"
const DiscountUrl = BaseUrl + "/gunun-firsati-teklifi"

type DiscountItem struct {
	imgUrl string
	url string
	name string
	currentPrice string
	oldPrice string // * default string
}

func hepsiburada(){
	client := &http.Client{}

	req, err := http.NewRequest("GET", DiscountUrl, nil)
	req.Header = http.Header{
		"user-agent":  []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.99 Safari/537.36"},
	}

	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	var products []DiscountItem

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)

		doc := soup.HTMLParse(bodyString)

		links := doc.FindAll("a", "class", "deal-of-the-day-item")
		for _, link := range links {
			item := DiscountItem{}

			item.url = BaseUrl + link.Attrs()["href"]
			item.name = link.Find("h3", "class", "deal-of-the-day-name").Text()
			item.imgUrl = link.Find("div", "class", "deal-of-the-day-image").Find("img").Attrs()["src"]
			item.currentPrice = link.Find("span", "class", "product-price").Text()

			if link.Find("del", "class", "product-old-price").Error == nil {
				item.oldPrice = link.Find("del", "class", "product-old-price").Text()
			}

			products = append(products,item)
		}
	}

	buildAndSendMessages(products)

}

// This is normal text - <b>and this is bold text</b>.\n<a href=\"https://www.carspecs.us/photos/c8447c97e355f462368178b3518367824a757327-2000.jpg\">test</a>
func buildAndSendMessages(products []DiscountItem) {
	for _, product := range products {
		message := fmt.Sprintf("<a href=\"%s\">%s</a>\n<del>%s</del>\n<b>%s</b>", product.url,product.name,product.oldPrice, product.currentPrice)
		//println(product.name)
		notify(message, product.imgUrl)
		break
	}
}
