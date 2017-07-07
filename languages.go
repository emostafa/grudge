package grudge

import (
	"fmt"
)

func Python(code string, testCases []string) (string, error) {
	filename := "pythonrunner.py"
	err := writeToFile("/tmp/", filename, code)
	if err != nil {
		return "", err
	}
	out, err := execCommand("python", "/tmp/"+filename)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", out), nil
}

func Javascript(code string, testCases []string) (string, error) {
	filename := "jsrunner.js"
	err := writeToFile("/tmp/", filename, code)
	if err != nil {
		return "", err
	}

	out, err := execCommand("node", filename)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", out), nil
}

func Ruby(code string, testCases []string) (string, error) {
	filename := "rubyrunner.rb"
	err := writeToFile("/tmp/", filename, code)

	out, err := execCommand("ruby", filename)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", out), nil
}
