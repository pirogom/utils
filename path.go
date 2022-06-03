package utils

import "golang.org/x/sys/windows"

/**
*	GetDesktopPath
**/
func GetDesktopPath() (string, error) {
	path, err := windows.KnownFolderPath(windows.FOLDERID_Desktop, 0)
	return path, err
}

/**
*	GetDocumentPath
**/
func GetDocumentPath() (string, error) {
	path, err := windows.KnownFolderPath(windows.FOLDERID_Documents, 0)
	return path, err
}
