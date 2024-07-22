package file_storage

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

var storageBasePath = "./public/storage/"

func Store(ctx *fiber.Ctx, file *multipart.FileHeader, dir string) (path string, err error) {
	if _, err = os.Stat(storageBasePath); os.IsNotExist(err) {
		if err = os.Mkdir(storageBasePath, os.ModePerm); err != nil {
			return "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}

	filePath := fmt.Sprintf("%s/%s-%s", dir, strconv.Itoa(int(time.Now().Unix())), file.Filename)

	storagePath := fmt.Sprintf("%s/%s", storageBasePath, filePath)

	if err = ctx.SaveFile(file, storagePath); err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return filePath, nil
}

func Remove(path string) (err error) {
	filePath := fmt.Sprintf("%s/%s", storageBasePath, path)

	if err = os.Remove(filePath); err != nil {
		if os.IsNotExist(err) {
			return fiber.NewError(fiber.StatusNotFound, "File Not Found")
		}

		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
