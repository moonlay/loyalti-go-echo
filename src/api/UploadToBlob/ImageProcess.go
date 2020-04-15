package UploadToBlob

import (
	"bytes"
	"fmt"
	"github.com/labstack/echo"
	"io"
	"os"
)

func Testku (c echo.Context) error {
	c.FormFile("upload")
	return nil
}

func ProcessImage(c echo.Context) error {
	fmt.Println("masuk sini")
	file, err := c.FormFile("upload")
	//files := c.FormValue("upload")
	//defer file.Close()
	if err != nil {
		fmt.Println("Error 1 : ", err.Error())
	}
	src, err := file.Open()
	if err != nil {
		fmt.Println("Error open  ",err.Error() )
	}

	dst, err := os.Create(file.Filename)
	fmt.Println("file : ", file.Filename)
	if err != nil {
		fmt.Println("Err dst : ", err.Error())
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		fmt.Println("Error 2 : ", err.Error())
	}
	fmt.Println("berhasil",dst)
	fmt.Println(buf.Bytes())
	UplodBytesToBlob(buf.Bytes())
	//fmt.Println("isi out : ", file)
	//fmt.Println("isi in : ", buf)
	//do other stuff
	return nil
}
