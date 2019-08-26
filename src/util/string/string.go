package string

func Count(arg1 string) (res1 int64) {
	if len(arg1) > 0 {
		return int64(len(arg1))
	}
	return 0
}
