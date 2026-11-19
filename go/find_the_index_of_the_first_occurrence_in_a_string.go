// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/

package main

func strStr(haystack string, needle string) int {
	for i, j := 0, 0; i < len(haystack); i++ {
		if haystack[i] == needle[j] {
			j++
		} else {
			i = i - j
			j = 0
		}

		if j == len(needle) {
			return i - j + 1
		}
	}

	return -1
}
