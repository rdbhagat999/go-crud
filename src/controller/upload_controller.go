package controller

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func saveFile(fileHeader *multipart.FileHeader, key int) error {
	fmt.Println(fileHeader.Filename)
	fileArray := strings.Split(fileHeader.Filename, ".")

	// if fileArray[1] != "png" {
	// 	return errors.New("pass a valid png")
	// }

	src, err := fileHeader.Open()

	if err != nil {
		return err
	}

	defer src.Close()

	// Create a new file in the desired destination folder
	dstPath := filepath.Join("./uploads", fileArray[0]+strconv.Itoa(key)+"."+fileArray[1])
	dst, err := os.Create(dstPath)

	if err != nil {
		return err
	}

	defer dst.Close()

	// fmt.Printf("File saved to: %s\n", dstPath)
	return nil

}

func UploadFile(ctx *gin.Context) {
	// form, err := ctx.MultipartForm()

	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// files := form.File["files"]

	// for key, file := range files {
	// 	err := saveFile(file, key)

	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// }

	// ctx.JSON(http.StatusOK, gin.H{"message": "Files uploaded successfuly"})

	name := ctx.PostForm("name")
	email := ctx.PostForm("email")

	// Multipart form
	form, frmErr := ctx.MultipartForm()

	if frmErr != nil {
		ctx.String(http.StatusBadRequest, "get form err: %s", frmErr.Error())
		return
	}
	files := form.File["files"]

	for _, file := range files {

		// Create a new file in the desired destination folder
		dstPath := filepath.Join("./uploads", file.Filename)

		if err := ctx.SaveUploadedFile(file, dstPath); err != nil {
			ctx.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}
	}

	ctx.String(http.StatusOK, "Uploaded successfuly %d files with fields name=%s and email=%s.", len(files), name, email)

}
