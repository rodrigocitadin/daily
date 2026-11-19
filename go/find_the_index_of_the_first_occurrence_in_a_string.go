// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/

package main

// import "strings"

// func strStr(haystack string, needle string) int {
// 	for i, j := 0, 0; i < len(haystack); i++ {
// 		if haystack[i] == needle[j] {
// 			j++
// 		} else {
// 			i = i - j
// 			j = 0
// 		}
//
// 		if j == len(needle) {
// 			return i - j + 1
// 		}
// 	}
//
// 	return -1
// }

// func strStr(haystack string, needle string) int {
// 	return strings.Index(haystack, needle)
// }

// func strStr(haystack string, needle string) int {
// 	for i := 0; i <= len(haystack)-len(needle); i++ {
// 		if haystack[i:len(needle)+i] == needle {
// 			return i
// 		}
// 	}
// 	return -1
// }

func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}

	lps := make([]int, m)
	for i, k := 1, 0; i < m; i++ {
		for k > 0 && needle[i] != needle[k] {
			k = lps[k-1]
		}

		if needle[i] == needle[k] {
			k++
		}

		lps[i] = k
	}

	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = lps[j-1]
		}

		if haystack[i] == needle[j] {
			j++
		}

		if j == m {
			return i - m + 1
		}
	}

	return -1
}
