package main

func canConstruct(ransomNote string, magazine string) bool {
	l1 := len(ransomNote)
	l2 := len(magazine)
	if l1 > l2 {
		return false
	}

	arr1 := []byte(ransomNote)
	arr2 := []byte(magazine)
	record := make(map[byte]int)
	for i := 0; i < l2; i++ {
		record[arr2[i]]++
		if i < l1 {
			record[arr1[i]]--
		}
	}

	for _, val := range record {
		if val < 0 {
			return false
		}
	}
	return true

}