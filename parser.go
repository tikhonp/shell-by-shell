package shellbyshell

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

const (
	shbyshToken = "shbysh"
	stepToken   = "step"
)

type Step struct {
	name string
	code string
}

func ParseFile(r io.Reader) ([]Step, error) {
	stringsBuffer := new(strings.Builder)
	_, err := io.Copy(stringsBuffer, r)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(stringsBuffer.String(), "\n")

	var steps []Step
	var currentStep Step

	for idx, line := range lines {
		if strings.Contains(line, shbyshToken) && currentStep.name == "" {
			lineTokens := strings.Split(line, shbyshToken)
			if len(lineTokens) < 2 {
				return nil, fmt.Errorf("invlid syntax on line %d", idx+1)
			}

			// check shbysh directive
			if strings.HasPrefix(lineTokens[1], stepToken) {
				lineTokens = strings.Split(line, ":")
				if len(lineTokens) < 2 {
					return nil, fmt.Errorf("invlid syntax on line %d", idx+1)
				}
				stepName := strings.TrimSpace(lineTokens[1])
				if len(stepName) == 0 {
					return nil, fmt.Errorf("empty step name on line %d", idx+1)
				}

			} else {
				return nil, fmt.Errorf("invalid shbysh directive on line %d", idx+1)
			}

		} else if currentStep.name != "" {

		}
	}

	if len(lines) < 2 {
		return nil, errors.New("found less than two lines in config script: please check the config file syntax")
	}
	if !strings.Contains(lines[0]) {

	}
	for idx := range len(lines) {

	}
}
