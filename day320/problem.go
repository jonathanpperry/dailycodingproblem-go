package day320

// SmallestWindowEveryDistinctLength returns the length
// of the smallest window that contains every distinct
// character available.
// Runs in O(N) time and O(k) extra space where k
// is the number of distinct characters.
func SmallestWindowEveryDistinctLength(str string) int {
	distinct := make(map[rune]int)
	for _, r := range str {
		distinct[r]++
	}
	distinctCount := len(distinct)
	distinct = make(map[rune]int, distinctCount)
	var begin int
	smallest := len(str)
	for end, r := range str {
		distinct[r]++
		if len(distinct) == distinctCount {
			for toRemove := rune(str[begin]); distinct[toRemove] > 1; toRemove = rune(str[begin]) {
				distinct[toRemove]--
				begin++
			}
			if length := end - begin + 1; length < smallest {
				smallest = length
			}
		}
		for len(distinct) == distinctCount {
			toRemove := rune(str[begin])
			distinct[toRemove]--
			if count := distinct[toRemove]; count == 0 {
				delete(distinct, toRemove)
			}
			begin++
		}
	}
	return smallest
}
