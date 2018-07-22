package problems

import "strings"

func isRotation(str1, str2 string) bool {
	// 1. check if length is same
	// 2. append str2 to itself so that we get a string in which
	//    all rotations are present as substring
	// 3. check if str1 is a substring in str2+str2
	return len(str1) == len(str2) && strings.Contains(str2+str2, str1)
}
