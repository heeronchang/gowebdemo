package httpclient

import (
	"context"
	"gowebdemo/configs/appone"
	"io"
	"net/http"
	"time"
)

func Request(url, method string, body io.Reader) ([]byte, error) {
	client := http.DefaultClient

	timeout := appone.GetTimeout()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return res, nil
}
