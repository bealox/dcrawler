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
	Name string //dog breed
	Link string
}

type Result struct {
	State   string
	Breeder []*Breeder
}

// TODO: Search function for matcher interface should return Result not []Result
type Matcher interface {
	Search(feed *Feed) ([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, results chan<- *Result) {
	matchResults, err := matcher.Search(feed)
	if err != nil {
		log.Println(err)
		return
	}

	for _, result := range matchResults {
		results <- result
	}
}

func Register(state string, matcher Matcher) {

	if _, exist := matchers[state]; exist {
		log.Fatalln(state, "Matcher already registered")
	}

	matchers[state] = matcher

}

func Run() {
	feeds, err := RetrieveFeed()

	if err != nil {
		log.Fatalln(err)
	}

	results := make(chan *Result)

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exists := matchers[feed.State]

		if !exists {
			matcher = matchers["default"]
		}

		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	go func() {
		waitGroup.Wait()
		close(results)
	}()

	Display(results)

}

func Display(results chan *Result) {
	for result := range results {
		if result.Breeder != nil {
			ProcessBreeder(result.Breeder)
		}
	}
}
