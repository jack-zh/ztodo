package utils

func RemoveSlice(slice []interface{}, start, end int) []interface{} {
	result := make([]interface{}, len(slice)-(end-start))
	at := copy(result, slice[:start])
	copy(result[at:], slice[end:])
	return result
}
