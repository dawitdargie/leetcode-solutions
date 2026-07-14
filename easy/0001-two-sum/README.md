# Two Sum

![Difficulty](https://img.shields.io/badge/Difficulty-Easy-green)

## Problem

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

 
Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].


Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]


Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]


 
Constraints:


	2 <= nums.length <= 104
	-109 <= nums[i] <= 109
	-109 <= target <= 109
	Only one valid answer exists.


 
Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?

## Solution

**Language:** Go  
**Runtime:** 4 ms (beats 35.62%)  
**Memory:** 5.8 MB (beats 68.38%)  
**Submitted:** 2026-07-14T15:08:00.904Z  

```go
func twoSum(nums []int, target int) []int {
    //I need place to remember numbers
    seen := make(map[int]int)
    //needed to look every number
    for index, num := range nums{
        //calculate the needed number
        needed := target - num
        //check if needed in the map
        i, exists := seen[needed]
        //if it exists
        if exists{
            return []int{i, index}
        }
        //if not exsisted, store it
        seen[num] = index
    }
    //if we don't get answer, return empty
    return []int{}
}
```

---

[View on LeetCode](https://leetcode.com/problems/two-sum/)