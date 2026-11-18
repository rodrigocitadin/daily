// https://leetcode.com/problems/find-the-difference/description/

package main

func findTheDifference(s string, t string) byte {
	saw := make(map[rune]int)
	for _, v := range s {
		saw[v]++
	}

	for _, v := range t {
		if saw[v] == 0 {
			return byte(v)
		}
		saw[v]--
	}

	return 0
}
