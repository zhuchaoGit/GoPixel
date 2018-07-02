package fileOrg

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"os"
	"log"
	"io"
	"github.com/zhuchao/GoPixel/lang"
)

func Upload(context *gin.Context) {
	file, header, err := context.Request.FormFile("image")
	if err != nil {
		context.JSON(http.StatusOK, jsonBody.UnKnownError(err))
		return
	}
	//文件的名称
	filename := header.Filename
	fmt.Println(file, err, filename)
	out, err := os.Create("asserts/img/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	context.JSON(http.StatusOK, jsonBody.Success())
}
