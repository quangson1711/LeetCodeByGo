package main

func convertToTitle(columnNumber int) string {
	columnName := ""
	for columnNumber > 0 {
		rem := columnNumber % 26
		if rem == 0 {
			columnName = columnName + "Z"
			columnNumber = (columnNumber / 26) - 1
		} else {
			columnName = columnName + ((string)(rune((rem - 1) + 'A')))
			columnNumber = columnNumber / 26
		}
	}
	return Reverse(columnName)
}
