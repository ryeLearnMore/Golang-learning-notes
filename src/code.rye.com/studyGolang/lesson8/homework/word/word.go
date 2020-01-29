package word

import "unicode"

// 跟文本处理相关的函数

// IsPalindrome 是一个回文判断函数，返回布尔值
// func IsPalindrome(s string) bool {
// 	s2 := []rune(s)
// 	lenStr := len(s2)
// 	for i := 0; i < lenStr/2; i++ {
// 		if s2[i] != s2[lenStr-i-1] {
// 			return false
// 		}
// 	}
// 	return true
// }

// IsPalindrome 是一个回文判断函数，返回布尔值
func IsPalindrome(s string) bool {
	// rune的常用用法：https://www.jianshu.com/p/4fbf529926ca
	var letters []rune
	for _, l := range s {
		// 判断l是不是一个letter
		if unicode.IsLetter(l) {
			// unicode.ToLower(l)转换成小写
			letters = append(letters, unicode.ToLower(l))
		}
	}

	for i := 0; i < len(letters)/2; i++ {
		if letters[i] != letters[len(letters)-i-1] {
			return false
		}
	}

	return true
}
