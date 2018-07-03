package fileOrg

func TemplatePath() string {
	return "asserts/html/index_0000000000.html"
}
func TemplateFile() string {
	return "index_0000000000.html"
}
func ImgPath() string {
	return "asserts/img"
}
func HtmlPath() string {
	return "asserts/html"
}
func HtmlOf(fileName string) string {
	return "asserts/html/" + fileName
}
