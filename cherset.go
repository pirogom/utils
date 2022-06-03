package utils

import (
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

/**
* cp949 to utf8 byte
**/
func AnsiToUTF8(src []byte) []byte {
	got, _, _ := transform.String(korean.EUCKR.NewDecoder(), string(src))

	return []byte(got)
}

/**
* cp949 to utf8 string
**/
func AnsiToUtf8String(src string) string {
	got, _, _ := transform.String(korean.EUCKR.NewDecoder(), src)

	return got
}

/**
*	utf8 to cp949
**/
func Utf8ToAnsi(src []byte) []byte {
	got, _, _ := transform.String(korean.EUCKR.NewEncoder(), string(src))

	return []byte(got)
}
