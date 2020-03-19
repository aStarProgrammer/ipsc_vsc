package main

import (
	"ipsc_vsc/Configuration"
	"ipsc_vsc/Utils"

	//"fmt"
	"io/ioutil"
)

const (
	QUICKHELP = "QUICKHELP"
	FULLHELP  = "FULLHELP"
)

func GetQuickHelpInformation() (string, error) {
	quickHelpFilePath, errPath := Configuration.GetQuickHelpPath()
	if errPath != nil {
		Utils.Logger.Println("GetQuickHelp: " + errPath.Error())
		return "", errPath
	}
	bHelpContent, errRead := ioutil.ReadFile(quickHelpFilePath)

	if errRead != nil {
		Utils.Logger.Println("GetQuickHelp: " + errRead.Error())
		return "", errRead
	}
	var sHelpContent = string(bHelpContent)
	return sHelpContent, nil
}

func GetFullHelpInformation() (string, error) {
	fullHelpFilePath, errPath := Configuration.GetFullHelpPath()

	if errPath != nil {
		Utils.Logger.Println("GetFullHelp: " + errPath.Error())
		return "", errPath
	}

	bHelpContent, errRead := ioutil.ReadFile(fullHelpFilePath)

	if errRead != nil {
		Utils.Logger.Println("GetFullHelp: " + errRead.Error())
		return "", errRead
	}
	var sHelpContent = string(bHelpContent)
	return sHelpContent, nil
}

func GetFieldHelpMsg(fieldName string) string {
	switch fieldName {
	case "Command":
		return "Command you want to run,for example, Compile:Compile the whole site, run ipsc -h or -Help to get more information. (If empty, Help command)"
	case "SiteTitle":
		return "Title of the site and site project,will show on the top of page of site, the shorter,the better"
	case "SiteDescription":
		return "Description of the site and site project, will show in the start section of index.html,the shorter the better"
	case "SiteAuthor":
		return "Author of site, if empty, will use current login user"
	case "SiteFolder":
		return "Path of Site Project "
	case "ExportFolderPath":
		return "The folder you want to export site source pages"
	case "OutputFolder":
		return "Folder used to put Compiled html files, if empty, will created automatically as sub folder Output of site project folder, if specified, will use the folder you specified"
	case "PropertyName":
		return "Name of property that you want to update,get or delete"
	case "PropertyValue":
		return "Value of property,enter it as string, IPSC will convert it to needed type"
	case "IndexPageSize":
		return "The size of index and more page, normal page contains 20 items,small page contains 10 items, very small page contains 5 items, big page contains 30 items (Default Normal)"
	case "StopCompileWithError":
		return "Whether stop Compile when encounter error during compling (Default true)"
	case "VerboseLog":
		return "Whether output verbose log of compling (Default false)"
	case "HelpType":
		return "FullHelp or QuickHelp (If empty, QuickHelp)"
	case "PageID":
		return "ID for page source file, you need to set this field when operating page source file"
	case "PageTitle":
		return "Title for page source file, will be used as page title, it should be less than 30 words"
	case "PagePath":
		return "Path of page source file"
	case "PageAuthor":
		return "Author of the page,if empty, will use current login user"
	case "IsTop":
		return "The page is on top of index page? (Default false)"
	case "PageType":
		return "Type of Page, Markdown,Html or Link, if empty (Default Markdown)"
	case "LinkUrl":
		return "Url when you user Link commands, AddLink UpdateLink DeleteLink GetLink"
	case "TitleImage":
		return "Image for page , will displayed with title of page in index page"
	case "RestorePage":
		return "Whether move page source file to recycled bin when you delete it (Default true)"
	case "MarkdownType":
		return "The type of markdown page source file you want to create, now supports Blank,News (Default News)"
	case "FilePath":
		return "The file you want to add, can be folder or file. Absolute File Path when you add file, relative path when you delelete file. You can get relative path of files by COMMAAND ListFile"
	case "AddFileForce":
		return "Whether replace file without confirm, if false, will confirm before replace a file that already exist"
	default:
		return ""
	}
}
