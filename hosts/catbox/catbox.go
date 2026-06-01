package catbox

import (
	"io"
	"main/utils"
	"os"
)

const (
	uploadUrl = "https://catbox.moe/user/api.php"
	referer   = "https://catbox.moe/"
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
	size, err := utils.CheckSize(path, "200MB")
	if err != nil {
		return "", err
	}
	userHash := os.Getenv("CATBOX_USERHASH")
	formMap := map[string]string{
		"userhash": userHash,
		"reqtype":  "fileupload",
	}
	headers := map[string]string{
		"Referer": referer,
	}
	fileUrl, err := upload(path, size, args.ByteLimit, formMap, headers)
	return fileUrl, err
}
