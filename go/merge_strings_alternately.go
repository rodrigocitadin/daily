// https://leetcode.com/problems/merge-strings-alternately/description/

package main

func mergeAlternately(word1 string, word2 string) string {
	var result []byte
	for i := range max(len(word1), len(word2)) {
		if i < len(word1) {
			result = append(result, word1[i])
		}

		if i < len(word2) {
			result = append(result, word2[i])
		}
	}

	return string(result)
}
