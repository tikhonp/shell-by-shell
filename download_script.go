package shellbyshell

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func Download(url string) (io.ReadCloser, error) {
    if strings.HasPrefix(url, "http") {
        resp, err := http.Get(url)
        return resp.Body, err
    } else {
        f, err := os.Open(url)
        return f, err
    }
}
