package pkg

func CheckWithCustomError(e error, message string) {
	if e != nil {
		panic(message)
	}
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
