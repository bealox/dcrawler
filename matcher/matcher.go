package matcher

import (
	"log"
	"sync"
)

var (
	matchers = make(map[string]Matcher)
)

// TODO: change structs in matcher
// Result : state and []Breed
// Breed: Name, Link, []Breeder
// Breeder Breed, ID, Email

type Breeder struct {
	Breed string
	ID    string
	Email string
}

type Breed struct {
	Name    string //dog breed
	Link    string
	Breeder []*Breeder
}

type Result struct {
	State string
	Breed []*Breed
}

type Matcher interface {
	Search(feed *Feed) (*Result, error)
}

func Match(matcher Matcher, feed *Feed) *Result {
	matchResults, err := matcher.Search(feed)
	if err != nil {
		// log.Println(err)
		return nil
	}

	return matchResults

}

func Register(state string, matcher Matcher) {

	if _, exist := matchers[state]; exist {
		log.SetPrefix("Error")
		log.Fatalln(state, "Matcher already registered")
	}

	matchers[state] = matcher

}

func Run() {
	feeds, err := RetrieveFeed()

	if err != nil {
		log.Println(err)
	}

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exists := matchers[feed.State]

		if !exists {
			matcher = matchers["default"]
		}

		go func(matcher Matcher, feed *Feed) {
			result := Match(matcher, feed)
			if result != nil {
				ProcessBreeder(result.Breed, result.State)
			}
			waitGroup.Done()
		}(matcher, feed)
	}

	waitGroup.Wait()

}
