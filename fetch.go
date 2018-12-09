package motest

import (
	"bytes"
	"net/http"

	"github.com/pkg/errors"
)

func fetch(req *crawlRequest) *crawlResponse {
	resp, err := http.Get(req.URL)
	if err != nil {
		return &crawlResponse{
			req: req,
			err: errors.Wrap(err, "http.Get() failed"),
		}
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return &crawlResponse{
			req:        req,
			statusCode: resp.StatusCode,
		}
	}

	buf := new(bytes.Buffer)
	nbytes, err := buf.ReadFrom(resp.Body)
	out := &crawlResponse{
		req:        req,
		statusCode: resp.StatusCode,
		body:       buf.Bytes(),
	}
	if err != nil {
		out.err = errors.Wrapf(
			err,
			"ReadFrom(resp.Body) failed, read: %v bytes, contentLength: %v",
			nbytes,
			resp.ContentLength,
		)
	}
	return out
}
