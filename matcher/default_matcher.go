package matcher

import ()

type defaultMatcher struct{}

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

func (m defaultMatcher) Search(feed *Feed) ([]*Result, error) {
	return nil, nil
}
