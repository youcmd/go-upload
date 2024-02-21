package litterbox

import (
	"io"
	"main/utils"
)

const (
	uploadUrl = "https://litterbox.catbox.moe/resources/internals/api.php"
	referer   = "https://litterbox.catbox.moe/"
)

func upload(path string, size, byteLimit int64, formMap, headers map[string]string) (string, error) {
	respBody, err := utils.MultipartUpload(uploadUrl, path, "fileToUpload", size, byteLimit, formMap, nil, headers)
	if err != nil {
		return "", err
	}
	defer respBody.Close()
	bodyBytes, err := io.ReadAll(respBody)
	return string(bodyBytes), err
}

func Run(args *utils.Args, path string) (string, error) {
	size, err := utils.CheckSize(path, "1000MB")
	if err != nil {
		return "", err
	}
	formMap := map[string]string{
		"reqtype":  "fileupload",
		"time":  "72h",
	}
	headers := map[string]string{
		"Referer": referer,
	}
	fileUrl, err := upload(path, size, args.ByteLimit, formMap, headers)
	return fileUrl, err
}
