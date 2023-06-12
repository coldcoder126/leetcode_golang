package main

import "fmt"

func main() {
	ans := reverseWords2("this is  my   name")
	fmt.Println(ans)
}

func reverseWords(s string) string {
	//1.使用双指针删除冗余的空格
	slowIndex, fastIndex := 0, 0
	b := []byte(s)
	//删除头部冗余空格
	for len(b) > 0 && fastIndex < len(b) && b[fastIndex] == ' ' {
		fastIndex++
	}
	//删除单词间冗余空格
	for ; fastIndex < len(b); fastIndex++ {
		if fastIndex-1 > 0 && b[fastIndex-1] == b[fastIndex] && b[fastIndex] == ' ' {
			continue
		}
		b[slowIndex] = b[fastIndex]
		slowIndex++
	}
	//删除尾部冗余空格
	if slowIndex-1 > 0 && b[slowIndex-1] == ' ' {
		b = b[:slowIndex-1]
	} else {
		b = b[:slowIndex]
	}
	//2.反转整个字符串
	reverse(&b, 0, len(b)-1)
	//3.反转单个单词  i单词开始位置，j单词结束位置
	i := 0
	for i < len(b) {
		j := i
		for ; j < len(b) && b[j] != ' '; j++ {
		}
		reverse(&b, i, j-1)
		i = j
		i++
	}
	return string(b)
}

func reverse(b *[]byte, left, right int) {
	for left < right {
		(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
		left++
		right--
	}
}

func reverseWords2(s string) string {
	b := []byte(s)
	l1, l2 := 0, len(b)-1
	//去除头部空格
	for b[l1] == ' ' {
		l1++
	}
	//去除尾部空格
	for l2 == ' ' {
		l2--
	}
	//更新数组
	b = b[l1 : l2+1]

	//删除中间多余空格
	l1, l2 = 0, 0
	for l2 < len(b)-1 {
		l1++
		l2++
		for l2 > 0 && b[l2] == b[l2-1] && b[l2] == ' ' {
			l2++
		}
		b[l1] = b[l2]
	}
	b = b[:l1+1]

	//将字符串反转
	l1, l2 = 0, len(b)-1
	reverse(&b, l1, l2)

	//再将每个单词反转
	l1, l2 = 0, 0
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			l2 = i - 1
			reverse(&b, l1, l2)
			l1 = i + 1
		} else if i == len(b)-1 {
			l2 = i
			reverse(&b, l1, l2)
		}
	}

	return string(b)
}
