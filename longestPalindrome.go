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
/*
对回文的理解又错了一次，尼玛，写成首尾字符对应的算法了，菜鸡伤不起啊
*/
func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	var j, length, maxlen, start, end int = 0, len(s), 0, 0, 0
	var set map[byte]int = make(map[byte]int)
	for j < length {
		i, ok := set[s[j]]
		if ok {
			if (j - i) > maxlen {
				maxlen, start, end = j-i, i, j+1
			}
		} else {
			set[s[j]] = j
		}
		j++
	}
	return s[start:end]
}

/**
* 正确理解题目后，按惯例来一个暴力算法测试，居然没有超时
*/
func isPalindrome(s string) bool {
	middle := (len(s)+1)/2 - len(s)%2
	for i := 0; i < middle; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
func longestPalindrome(s string) string {
	max, start, end := 0, 0, 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isPalindrome(s[i:(j + 1)]) {
				if (j-i)+1 > max {
					max, start, end = (j - i), i, (j + 1)

				}
			}
		}
	}
	return s[start:end]
}

/*
看了评论，用了动态规划DP
*/
func longestPalindrome(s string) string {
	var n, start, end, max int = len(s), 0, 0, 0
	var dp [1000][1000]bool
	for i := 0; i < n; i++ {
		for j := i; j >= 0; j-- {
			if s[i] == s[j] && ((i-j) < 3 || dp[i-1][j+1]) {
				dp[i][j] = true
			}
			if dp[i][j] && (i-j+1) > max {
				max = i - j + 1
				start = j
				end = i + 1
			}
		}
	}
	return s[start:end]
}


