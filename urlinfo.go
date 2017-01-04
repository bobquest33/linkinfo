package linkinfo

import (
	"net/http"
	"strings"
)

func getUrlMimeType(url string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return "", nil
	}
	if resp, err := client.Do(req); err != nil {
		return "", err
	} else {
		defer resp.Body.Close()
		if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			return strings.Trim(strings.Split(resp.Header.Get("Content-Type"), ";")[0], " "), nil
		}
	}
	return "", nil
}
