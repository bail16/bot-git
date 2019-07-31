package newsSrc

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/blacklists"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/messages"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/newsSrc/newsAbstract"
)

var GetVoyage = []newsAbstract.GetNews{
	voyageSpider,
	voyageMlecznePodroze,
}

var VoyagePage = map[string]int{
	"Spider": 0,
	"MlecznePodroze": 0,
}
func voyageSpider() []messages.Message{
	blacklists.New("voyageSpiderBL")
	VoyagePage["Spider"]++
	return newsAbstract.GetSpider("podroze", VoyagePage["Spider"])
}

func voyageMlecznePodroze() []messages.Message{
	blacklists.New("voyageMlecznePodrozeBL")
	VoyagePage["MlecznePodroze"]++
	doc := abstract.GetDoc(fmt.Sprintf("https://mlecznepodroze.pl/tag/news/page/%v/", VoyagePage["MlecznePodroze"]))

	div := abstract.GetDiv(doc,"div.primary-post-content")

	var news []messages.Message

	div.Each(func(i int, s *goquery.Selection){

		image,_ := s.Find("div.picture > div.picture-content > a > img").Attr("src")
		text,_ := s.Find("div.picture > div.picture-content > a").Attr("title")
		textlink, _ := s.Find("div.picture > div.picture-content > a").Attr("href")
		temp := messages.Message{
			TitleLink:  textlink,
			Img: messages.Image{
				Header: text,
				ImageUrl: image,
			},
		}

		if !temp.Img.IsEmpty() && temp.TitleLink != ""{
			news = append(news,temp)
		}
	})
	return news
}