package problems

//Write a method to replace all spaces in a string with ‘%20’.
func replaceSpace(str1 string) string {
	modifiedString := ""
	for _, char := range str1 {
		if char == ' ' {
			modifiedString += "%20"
			continue
		}
		modifiedString += string(char)
	}
	return modifiedString
}
