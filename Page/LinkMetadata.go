package Page

import (
	"encoding/json"
	"errors"
	"ipsc_vsc/Utils"
	"os"
)

type LinkMetadata struct {
	Url    string
	Title  string
	Author string
	IsTop  bool
}

func LinkMeta2Json(link LinkMetadata) (string, error) {

	_jsonbyte, errJson := json.Marshal(link)

	if errJson != nil {
		return "", errJson
	}
	return string(_jsonbyte), nil
}

func SaveLinkMetadataToFile(filePath string, link LinkMetadata) (bool, error) {
	if "" == filePath {
		return false, errors.New("LinkMetadata.SaveLinkMetadataToFile:FilePath is empty")
	}

	linkStr, errConvert := LinkMeta2Json(link)

	if errConvert != nil {
		return false, errors.New("LinkMetadata.SaveLinkMetadataToFile:cannot convert LinkProperties to json string")
	}

	var errFilePath error
	if !Utils.PathIsExist(filePath) {
		filePath, errFilePath = Utils.MakePath(filePath)
		if errFilePath != nil {
			return false, errors.New("LinkMetadata.SaveLinkMetadataToFile:Path nor exist and create parent folder failed")
		}
	}
	//路径分为绝对路径和相对路径
	//create，文件存在则会覆盖原始内容（其实就相当于清空），不存在则创建
	fp, error := os.Create(filePath)
	if error != nil {
		return false, error
	}
	//延迟调用，关闭文件
	defer fp.Close()

	_, errWriteFile := fp.WriteString(linkStr)

	if errWriteFile != nil {
		return false, errors.New("LinkMetadata.SaveLinkMetadataToFile:Write json to file failed")
	}

	return true, nil
}
