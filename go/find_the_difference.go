// https://leetcode.com/problems/find-the-difference/description/

package main

// func findTheDifference(s string, t string) byte {
// 	saw := make(map[rune]int)
// 	for _, v := range s {
// 		saw[v]++
// 	}
//
// 	for _, v := range t {
// 		if saw[v] == 0 {
// 			return byte(v)
// 		}
// 		saw[v]--
// 	}
//
// 	return 0
// }

func findTheDifference(s string, t string) byte {
	var diff byte

	for i := 0; i < len(s); i++ {
		diff ^= s[i]
	}

	for i := 0; i < len(t); i++ {
		diff ^= t[i]
	}

	return diff
}
