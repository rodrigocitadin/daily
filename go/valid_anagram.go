// https://leetcode.com/problems/valid-anagram/description/

package main

import (
	"sync"
	"sync/atomic"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var counts [256]atomic.Int64
	var wg sync.WaitGroup

	wg.Go(func() {
		for _, v := range s {
			counts[v].Add(1)
		}
	})

	wg.Go(func() {
		for _, v := range t {
			counts[v].Add(-1)
		}
	})

	wg.Wait()

	for i := range counts {
		if counts[i].Load() != 0 {
			return false
		}
	}

	return true
}
