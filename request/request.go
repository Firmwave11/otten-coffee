package request

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

func Curl(ctx context.Context, url, method string, timer time.Duration, header map[string]string, payload *bytes.Buffer) (context.Context, []byte, int, error) {
	var req *http.Request

	if method == http.MethodGet {
		request, err := http.NewRequest(method, url, nil)
		if err != nil {
			return ctx, nil, http.StatusInternalServerError, err
		}
		req = request
	} else {
		request, err := http.NewRequest(method, url, payload)
		if err != nil {
			return ctx, nil, http.StatusInternalServerError, err
		}
		req = request
	}

	client := http.Client{
		Timeout: timer,
	}

	for key, val := range header {
		req.Header.Set(key, val)
	}

	res, err := client.Do(req)
	if err != nil {
		return ctx, nil, http.StatusInternalServerError, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ctx, nil, http.StatusInternalServerError, err
	}

	return ctx, body, res.StatusCode, nil
}
