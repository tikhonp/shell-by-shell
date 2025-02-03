package shellbyshell

import (
	"errors"
	"flag"
)

type FlagConfig struct {
	Url string
}

func ParseFlags() (*FlagConfig, error) {
	url := flag.String("url", "", "file or http url to config script")
	flag.Parse()
	if *url == "" {
		return nil, errors.New("provide --url")
	}
	return &FlagConfig{
		Url: *url,
	}, nil
}
