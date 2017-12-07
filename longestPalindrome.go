/*
第一次又理解错题目了，可能是英文理解实在太差了，事实上又犯了看题不仔细的毛病
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
Example:
Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example:
Input: "cbbd"
Output: "bb"

我看了题目之后，马上写了一个把字符串中重复字符最长的算法出来,相对于前几天，编码的速度快了很多，事实上，题目是“最长回路”，结果当然是错的
*/
func longestPalindrome(s string) string {
	var i, j, length, maxlen int = 0, 1, len(s), 1
	var start, end = 0, 1
	for j < length {
		if s[j] != s[i] {
			i = j
			j = i + 1
		} else {
			j++
		}
		if (j - i) > maxlen {
			maxlen = j - i
			start, end = i, j
		}
	}
	return s[start:end]
}




