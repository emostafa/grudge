package grudge

import (
	"fmt"
	"os"
	"os/exec"
)

func writeToFile(dir string, filename string, data string) error {
	f, err := os.Create(dir + filename)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(data)

	return nil
}

func execCommand(language string, filename string) (string, error) {
	bin := ""
	switch language {
	case "javascript":
		bin = "node"
	default:
		bin = language
	}
	cmd := exec.Command(bin, filename)
	cmd.Dir = "/tmp"
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return fmt.Sprintf("%s", out), nil
}
