package matcher

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

type qldMatcher struct{}

const (
	uri string = "http://www.dogsqueensland.org.au"
)

func init() {
	var matcher qldMatcher
	Register("QLD", matcher)
}

func (m qldMatcher) Search(feed *Feed) ([]*Result, error) {
	doc, err := goquery.NewDocument(feed.Link)

	var breeds []*Breed
	var results []*Result

	if err != nil {
		log.Fatal(err)
	}

	//Find breeds on this page
	doc.Find(".col-xs-7").Each(func(i int, s *goquery.Selection) {
		anchor := s.Find("a")
		href, exist := anchor.Attr("href")
		breed := anchor.Find("h5").Text()
		if !exist {
			return
		}

		breeds = append(breeds, &Breed{
			Name: breed,
			Link: href,
		})

	})

	breeders, errBreeder := RetrieveBreeder(breeds)

	if errBreeder != nil {
		return nil, errBreeder
	}

	results = append(results, &Result{
		State:   "QLD",
		Breeder: breeders,
	})

	return results, nil

}

func RetrieveBreeder(breeds []*Breed) ([]*Breeder, error) {

	log.Printf("How many Breed in QLD %d", len(breeds))

	var breeders []*Breeder

	for _, breed := range breeds {
		//Use GoRountine to run this process cocurrently
		// FIXME: use GORoutine to run goquery (maybe quicker)
		doc, err := goquery.NewDocument(uri + breed.Link)

		if err != nil {
			return nil, err
		}

		doc.Find(".col-xs-8").Each(func(i int, s *goquery.Selection) {
			id := strings.TrimSpace(s.Find("h3").Text())

			// TODO: maybe store webstie address in the database
			// QLD council put email and website under the same div
			// splitting out new line character, you can also get the website address via
			// fitleredEmail[2]
			email := strings.TrimSpace(s.Find(".email-content").Text())
			filteredEmail := strings.SplitN(email, "\n", -1)

			breeders = append(breeders, &Breeder{
				Breed: breed.Name,
				ID:    id,
				Email: filteredEmail[0],
			})

		})

	}

	log.Println("breeders count ")
	log.Println(len(breeders))

	return breeders, nil
}
