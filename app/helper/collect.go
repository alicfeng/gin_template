package helper

// SliceContain 判断切片是否包含元素 /**
func SliceContain(slice []string, value string) (int, bool) {
	for index, item := range slice {
		if value == item {
			return index, true
		}
	}

	return 0, false
}
