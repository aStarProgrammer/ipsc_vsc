package main

import (
	"fmt"
	"ipsc_vsc/Utils"
)

func testCopyFolder() {
	var src = "D:\\softwares"
	var dst = "F:\\Dst"
	var addForce = true

	Utils.CopyFolder(src, dst, addForce)
}

func testGetCommandArgs() {
	argList := GetUpdateCommandArgs()

	for _, arg := range argList {
		fmt.Println(arg)
	}
}

func testParseUpdateCommandArgs() {
	var cp CommandParser
	cp.ParseCommand()

	fmt.Println(cp.LinkUrl)
	fmt.Println(cp.PageTitle)
	fmt.Println(cp.PageAuthor)
	fmt.Println(cp.PageTitleImagePath)

	argList := GetUpdateCommandArgs()

	cp.ParseUpdateCommandArgs(argList)

	fmt.Println(cp.LinkUrl)
	fmt.Println(cp.PageTitle)
	fmt.Println(cp.PageAuthor)
	fmt.Println(cp.PageTitleImagePath)
}

func test() {
	testParseUpdateCommandArgs()
}
