/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

*/
//改了好几次，改少了很多代码，请原谅我的强迫症 */
func twoSum(nums []int, target int) []int {
	var tracker map[int]int = make(map[int]int, len(nums))
	for i, v := range nums {
		tv, ok := tracker[v]
		if ok {
            return []int{tv,i}
		}
        tracker[target-v] = i
	}
    return []int{}
}
