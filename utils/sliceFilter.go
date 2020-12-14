package utils

// FilterSliceString retorna uma nova Slice contendo apenas os elementos que passarem no filtro especificado
func FilterSliceString(s []string, testFunc func(element string) bool) []string {
	var newSlice []string
	for _, element := range s {
		if testFunc(element) {
			newSlice = append(newSlice, element)
		}
	}
	return newSlice
}
