package main

func main() {

}

func majorityElement(nums []int) int {
	mapCount := map[int]int{}
	for _, num := range nums {
		_, ok := mapCount[num]
		if !ok {
			mapCount[num] = 1
		} else {
			mapCount[num] += 1
		}
		if mapCount[num] > len(nums)/2 {
			return num
		}
	}
	return 0
}
