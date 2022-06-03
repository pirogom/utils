package utils

import (
	"io"
	"os"
)

/**
*	CopyFile
**/
func CopyFile(srcFile string, destFile string, deleteExist bool) error {
	// 기존 파일이 있으면 삭제
	if deleteExist && IsExistFile(destFile) {
		os.Remove(destFile)
	}
	//

	from, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer from.Close()

	to, err := os.OpenFile(destFile, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, from)

	if err != nil {
		return err
	}
	return nil
}

/**
* Exist check .. file or directory
**/
func IsExistFile(fname string) bool {
	if _, err := os.Stat(fname); os.IsNotExist(err) {
		return false
	}
	return true
}

/**
*	IsExistFileWithSize
**/
func IsExistFileWithSize(fname string) bool {

	fi, err := os.Stat(fname)

	if os.IsNotExist(err) {
		return false
	}

	if fi.Size() > 0 {
		return true
	}

	return false
}
