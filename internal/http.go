package internal

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

type RequestCallback func(req *http.Request)

var httpClient = &http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     false,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		ResponseHeaderTimeout: 15 * time.Second,
	},
}

func RequestJSONFromURL[T any](method string, url string, data io.Reader, callback RequestCallback) (*T, error) {
	bytes, err := RequestBytesFromURL(method, url, data, callback)
	if err != nil {
		return nil, err
	}

	return FromJSON[T](bytes)
}

func RequestBytesFromURL(method string, url string, data io.Reader, callback RequestCallback) ([]byte, error) {
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	if callback != nil {
		callback(req)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request returned non-successful status code: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
