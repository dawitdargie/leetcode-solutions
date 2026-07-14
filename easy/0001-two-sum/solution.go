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