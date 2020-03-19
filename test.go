package main

import (
	"ipsc_vsc/Utils"
)

func testCopyFolder() {
	var src = "D:\\softwares"
	var dst = "F:\\Dst"
	var addForce = true

	Utils.CopyFolder(src, dst, addForce)
}

func test() {
	testCopyFolder()
}
