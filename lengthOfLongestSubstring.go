/*

先来一波brute force，结果被吊打~~~
982 / 983 test cases passed.
Status: Time Limit Exceeded
Submitted: 0 minutes ago
Last executed input:
*/

func lengthOfLongestSubstring(s string) int {
    end := len(s)
	maxlen := 1
    if (len(s) == 0){
        return 0
    }
	for end > 0 {
		s = s[:end]
		start := 0
		for start < end-1 {
			ss := s[start:end]
			for ii, cc := range ss {
				ii++
				for _, ccc := range ss[ii:] {
					if cc == ccc {
						goto BK_LABEL
					}
				}
			}
			if (end - start) > maxlen {
				maxlen = end - start
			}
		BK_LABEL:
			start++
		}
		end--
	}
	return maxlen
}
