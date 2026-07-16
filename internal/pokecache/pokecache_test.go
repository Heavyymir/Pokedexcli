package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestGetAdd(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val	[]byte
	}{
		{
			key: 	"https://awebsite.com",
			val: 	[]byte("testdata"),
		},
		{
			key:	"https:bwebsite.com",
			val:	[]byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T){
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("Expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 10 * time.Millisecond
	const waitTime = 25 * time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok{
		t.Errorf("expected to not find key")
		return
	}
}
