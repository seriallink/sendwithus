package swu

func NotIn(value interface{}, list ...interface{}) bool {

	for _, element := range list {
		if value == element {
			return false
		}
	}

	return true

}
