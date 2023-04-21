package utils

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	userAgentValue = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36"
)

func ChangeUserAgent(ua string) {
	userAgentValue = ua
}

/**
*	GetHttpData
**/
func GetHttpData(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("user-agent", userAgentValue)

	client := &http.Client{}
	rep, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		rep.Body.Close()
	}()

	buf, bufErr := ioutil.ReadAll(rep.Body)
	if bufErr != nil {
		return nil, bufErr
	}

	return buf, nil
}

/**
* HTTP 파일 다운로드
**/
func DownloadFile(filepath string, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Request 객체 생성
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	//필요시 헤더 추가 가능
	req.Header.Add("user-agent", UserAgentValue)

	// Client객체에서 Request 실행
	client := &http.Client{}
	rep, err := client.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		rep.Body.Close()
	}()

	// Write the body to file
	_, err = io.Copy(out, rep.Body)
	if err != nil {
		return err
	}

	return nil
}
