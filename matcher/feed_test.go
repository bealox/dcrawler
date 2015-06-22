package matcher

import (
	"testing"
)

/*
 Test Retrieve Feeds from /data/council.json file
 Making sure, the system is getting the right data back.
*/
func TestRetrieveFeeds(t *testing.T) {
	feeds, err := RetrieveFeed()

	if err != nil {
		t.Error(err)
	}

	answers := []struct {
		link string
	}{
		{"http://www.dogsqueensland.org.au/Frontend/Breeder?Length=7"},
		{"http://www.dogsnsw.org.au/puppies/breeders-directory.html"},
		{"http://www.dogsvictoria.org.au/DogsPuppies/BuyingAPuppy/Breedersdirectory.aspx"},
		{"http://www.dogssa.com.au/?page_id=1374"},
		{"http://www.dogswest.com/dogswest/Breeders_Search.htm"},
		{"http://www.dogsact.org.au/Breeders.htm"},
		{"http://www.dogsnt.com.au/breeds-breeders/breeders-directory/"},
	}

	for i := 0; i < len(feeds); i++ {
		if feeds[i].Link != answers[i].link {
			t.Errorf("%s is %s and not %s", feeds[i].Site, feeds[i].Link, answers[i].link)
		}

	}

}
