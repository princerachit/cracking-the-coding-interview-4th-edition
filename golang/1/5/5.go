package problems

import "strings"

//Write a method to replace all spaces in a string with ‘%20’.
func replaceSpace(str1 string) string {
	return strings.Replace(str1, " ", "%20", -1)
}
