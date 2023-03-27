package main

func isMatch(s string, p string) bool {
	//1.先将两个字符串变为数组
	arr_s := []byte(s)
	arr_p := []byte(p)
	l := len(arr_s)

	for idx_s, idx_p := 0, 0; idx_s < l; idx_s++ {
		if arr_p[idx_p] == 'a' && arr_s[idx_s] == arr_p[idx_p] {
			idx_p++
			continue
		}
		if arr_p[idx_p] == byte('.') {

		}
	}
	return true
}
