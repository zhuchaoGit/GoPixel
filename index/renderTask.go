package index

import (
	"time"
	"log"
	"fmt"
	"os"
	"bufio"
	"io/ioutil"
	"html/template"
	"strings"
	"sort"
	"github.com/zhuchao/GoPixel/fileOrg"
)

func extraDirName(dir []os.FileInfo) []string {
	imgList := make([] string, len(dir), len(dir))
	id := 0
	for _, fileInfo := range dir {
		if strings.HasPrefix(fileInfo.Name(), ".") || strings.EqualFold(fileInfo.Name(), "") {
			continue
		}
		imgList[id] = fileInfo.Name()
		id++
	}
	return imgList[:id]
}
func Init() *[]byte {
	dir, _ := ioutil.ReadDir(fileOrg.HtmlPath())
	fileList := extraDirName(dir);
	sort.Strings(fileList)
	fileName := fileList[len(fileList)-1]
	fileContent, _ := ioutil.ReadFile(fileName);
	return &fileContent
}
func RenderTask(ptr **[]byte) {
	ticker := time.NewTicker(5 * time.Second)
	for t := range ticker.C {
		log.Printf("render new template:%d", t.Unix())
		fileName := fmt.Sprintf("index_%10d.html", t.Unix())
		file, outError := os.Create(fileOrg.HtmlOf(fileName))
		if outError != nil {
			log.Fatal(outError)
			continue
		}
		writer := bufio.NewWriter(file)
		tp, err := template.ParseFiles(fileOrg.HtmlOf("index_0000000000.html"))
		if err != nil {
			log.Fatal(err)
			continue
		}
		dir, _ := ioutil.ReadDir(fileOrg.ImgPath())
		imgIdList := extraDirName(dir)
		renderError := tp.Execute(writer, imgIdList);
		if renderError != nil {
			log.Fatal(renderError)
		}
		writer.Flush()
		file.Close()
		data, _ := ioutil.ReadFile(fileOrg.HtmlOf(fileName))
		*ptr = &data
	}
}
