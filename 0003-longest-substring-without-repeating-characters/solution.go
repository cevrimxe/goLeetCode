// O(n) solution

func lengthOfLongestSubstring(s string) int {
	lastSeen := make(map[byte]int)
	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		char := s[right]

		if index, ok := lastSeen[char]; ok && index >= left {
			left = index + 1
		}

		lastSeen[char] = right

		currentLen := right - left + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}
