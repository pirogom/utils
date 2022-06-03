package utils

import (
	"io/ioutil"
	"net/http"
)

var (
	UserAgentValue = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36"
)

/**
*	GetHttpData
**/
func GetHttpData(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("user-agent", UserAgentValue)

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
