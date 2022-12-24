package mycd

import (
	"fmt"
	"os"
	"strings"
)

type MyCd struct {
	CurrentDir string
	NewDir     string
	resultDir  string
}

func NewMyCd(CurrentDir string, NewDir string) *MyCd {
	resultDir := formatDir(CurrentDir, NewDir)
	return &MyCd{CurrentDir, NewDir, resultDir}
}

func formatDir(CurrentDir string, NewDir string) string {
	var t string
	if NewDir[0] == '/' {
		t = NewDir
	} else {
		t = CurrentDir + "/" + NewDir
	}
	dirs := strings.Split(t, "/")
	var dirList []string
	for _, dir := range dirs {
		if dir == "" || dir == "." {
			continue
		}
		if dir == ".." {
			if len(dirList) > 0 {
				dirList = dirList[:len(dirList)-1]
			}
			continue
		}
		dirList = append(dirList, dir)
	}
	var result string
	if len(dirList) == 0 {
		result = "/"
	}
	for _, dir := range dirList {
		result = result + "/" + dir
	}

	return result
}

func (mycd *MyCd) PrintDir() {
	if mycd.isValidDir() {
		fmt.Println(mycd.resultDir)
	} else {
		fmt.Println(mycd.NewDir + ": No such file or directory")
	}
}

func (mycd *MyCd) isValidDir() bool {
	_, err := os.Stat(mycd.resultDir)
	if err == nil {
		return true
	} else {
		return false
	}
}
