package matcher

import (
	"io/ioutil"
	"net/http"
)

func init() {
	var matcher nswMatcher
	Register("NSW", matcher)
}

type nswMatcher struct{}

func (m nswMatcher) Search(feed *Feed) (*Result, error) {

	// var results *Result

	resp, err := http.Get(feed.Link)
	if err != nil {
		return nil, err
	}

	_, err2 := ioutil.ReadAll(resp.Body)

	if err2 != nil {
		return nil, err
	}

	// results = append(results, &Result{
	// 	State: "NSW",
	// })

	return nil, nil
}
