package grudge

import (
	"fmt"
)

// Python runs python code in a secured environment
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

// Javascript runs javascript code inside a secured environment
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

// Ruby runs ruby code inside a secured environment
func Ruby(code string, testCases []string) (string, error) {
	filename := "rubyrunner.rb"
	err := writeToFile("/tmp/", filename, code)

	out, err := execCommand("ruby", filename)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", out), nil
}
