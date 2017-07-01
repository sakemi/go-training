package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type Func struct {
	f    func(string) (interface{}, error)
	done chan struct{}
}

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

func New(f Func) *Memo {
	return &Memo{Func: f, cache: make(map[string]*entry)}
}

type Memo struct {
	Func
	mu    sync.Mutex
	cache map[string]*entry
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		if canceled(memo.done) {
			e.res.value, e.res.err = nil, fmt.Errorf("Operation has been canceled.")
			delete(memo.cache, key)
		} else {
			e.res.value, e.res.err = memo.f(key)
		}

		close(e.ready)
	} else {
		memo.mu.Unlock()
		<-e.ready
	}
	return e.res.value, e.res.err
}

func canceled(done <-chan struct{}) bool {
	if done == nil {
		return false
	}
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	done := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	memo := New(Func{httpGetBody, done})
	urls := []string{
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"http://gopl.io",
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"http://gopl.io",
	}

	getConcurrent(urls, memo)
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func getConcurrent(URLs []string, m *Memo) {
	var n sync.WaitGroup
	for _, url := range URLs {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	time.Sleep(10 * time.Second)
	for _, url := range URLs {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}
