package middleware

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Upload_File(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("input-image")

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		src, err := file.Open()

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		defer src.Close()

		temp_file, err := ioutil.TempFile("uploads", "image-*.png")

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		defer temp_file.Close()

		writtenCopy, err := io.Copy(temp_file, src)

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		fmt.Println("written copy :", writtenCopy)

		data := temp_file.Name()
		filename := data[8:]

		c.Set("dataFile", filename)

		return next(c)
	}

}
