package fetch

import (
	"errors"
	"io"
	"net/http"
	"time"
)

var headers = map[string]string{
	"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36 Edg/119.0.0.0",
	"Accept":          "application/json, text/plain, */*",
	"Accept-Encoding": "gzip, deflate, br",
	"Content-type":    "application/x-www-form-urlencoded",
}

func commonClient() *http.Client {
	trans := &http.Transport{
		MaxConnsPerHost: 10,
		MaxIdleConns:    90,
	}
	client := http.Client{
		Timeout:   10 * time.Second,
		Transport: trans,
	}
	return &client
}

func addHeaders(req *http.Request) {
	for k, v := range headers {
		req.Header.Add(k, v)
	}
}

func Do(targetUrl string) ([]byte, error) {
	c := commonClient()
	req, err := http.NewRequest("GET", targetUrl, nil)
	if err != nil {
		return nil, err
	}
	addHeaders(req)
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("http response status code is not 200")
	}
	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	return bodyByte, err
}
