package jsonBody

type Json map[string]interface {
}

func UnKnownError(err error) Json {
	return Json{
		"code":    -1,
		"message": err.Error(),
	}
}
func Success() Json {
	return Msuccess("")
}
func Msuccess(str string) Json {
	return Json{
		"code":    0,
		"message": str,
	}
}
