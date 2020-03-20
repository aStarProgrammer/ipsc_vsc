package Page

import (
	"encoding/json"
	"errors"
	"ipsc_vsc/Utils"
	"os"
)

type PsfMetadata struct {
	Title  string
	Author string
	IsTop  bool
}

func PsfMetadata2Json(link PsfMetadata) (string, error) {

	_jsonbyte, errJson := json.Marshal(link)

	if errJson != nil {
		return "", errJson
	}
	return string(_jsonbyte), nil
}

func SavePsfMetadataToFile(filePath string, link PsfMetadata) (bool, error) {
	if "" == filePath {
		return false, errors.New("PsfMetadata.SavePsfMetadataToFile:FilePath is empty")
	}

	linkStr, errConvert := PsfMetadata2Json(link)

	if errConvert != nil {
		return false, errors.New("PsfMetadata.SavePsfMetadataToFile:cannot convert LinkProperties to json string")
	}

	var errFilePath error
	if !Utils.PathIsExist(filePath) {
		filePath, errFilePath = Utils.MakePath(filePath)
		if errFilePath != nil {
			return false, errors.New("PsfMetadata.SavePsfMetadataToFile:Path nor exist and create parent folder failed")
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
		return false, errors.New("PsfMetadata.SavePsfMetadataToFile:Write json to file failed")
	}

	return true, nil
}
