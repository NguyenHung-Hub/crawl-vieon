package util

func RemoveDuplicates(slice []string) []string {
	seen := make(map[string]bool)
	result := []string{}

	for _, val := range slice {
		if _, ok := seen[val]; !ok {
			seen[val] = true
			result = append(result, val)
		}
	}
	return result
}

func FilterStr(inputs []string, excludes []string) []string {

	filtered_list := []string{}

	for _, id := range inputs {
		if !Contains(excludes, id) {
			filtered_list = append(filtered_list, id)
		}
	}
	return filtered_list
}

func Contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// find items in slice2 that are not present in slice1
func FindItemNotInSlice(slice1, slice2 []string) []string {
	existingItems := make(map[string]bool)
	for _, item := range slice1 {
		existingItems[item] = true
	}

	var itemsNotInSlice1 []string
	for _, item := range slice2 {
		if !existingItems[item] {
			itemsNotInSlice1 = append(itemsNotInSlice1, item)
		}
	}

	return itemsNotInSlice1
}
