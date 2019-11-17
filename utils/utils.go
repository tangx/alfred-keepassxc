package utils

func IsError(err error) {
	if err != nil {
		panic(err)
	}
}
