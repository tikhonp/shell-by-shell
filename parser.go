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

func ParseFile(r io.ReadCloser) ([]Step, error) {
	stringsBuffer := new(strings.Builder)
	defer r.Close()
	_, err := io.Copy(stringsBuffer, r)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(stringsBuffer.String(), "\n")

	var steps []Step
	var currentStep Step

	for idx, line := range lines {
		if strings.Contains(line, shbyshToken) {
			if currentStep.name != "" {
				steps = append(steps, currentStep)
				currentStep = Step{}
			}
			lineTokens := strings.Split(line, shbyshToken)
			if len(lineTokens) < 2 {
				return nil, fmt.Errorf("invlid syntax on line %d", idx+1)
			}

			// check shbysh directive
			if strings.HasPrefix(strings.TrimSpace(lineTokens[1]), stepToken) {
				lineTokens = strings.Split(line, ":")
				if len(lineTokens) < 2 {
					return nil, fmt.Errorf("invlid syntax on line %d", idx+1)
				}
				stepName := strings.TrimSpace(lineTokens[1])
				if len(stepName) == 0 {
					return nil, fmt.Errorf("empty step name on line %d", idx+1)
				}

				currentStep.name = stepName
			} else {
				return nil, fmt.Errorf("invalid shbysh directive on line %d", idx+1)
			}

		} else if currentStep.name != "" {
			currentStep.code += line + "\n"
		}
	}
	if currentStep.name != "" {
		steps = append(steps, currentStep)
	}

	if len(steps) == 0 {
		return nil, errors.New("no shbysh directive in script")
	}

	return steps, nil
}
