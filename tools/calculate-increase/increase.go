package increase

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

// CalculateIncrease calculates the increase in disc size and memory usage when using the burner package
func CalculateIncrease() (string, string, error) {
	withBurnerSize, withBurnerMem, err := testMain("tools/calculate-increase/test-mains/withburner")
	if err != nil {
		return "", "", err
	}

	withoutBurnerSize, withoutBurnerMem, err := testMain("tools/calculate-increase/test-mains/withoutburner")
	if err != nil {
		return "", "", err
	}

	return formatSize(withBurnerSize - withoutBurnerSize), formatSize(withBurnerMem - withoutBurnerMem), nil
}

func formatSize(size int64) string {
	return fmt.Sprintf("%.2f Mb", float64(size)/1024.0/1024.0)
}

const testRuns = 10

func testMain(dir string) (int64, int64, error) {
	base, err := os.Getwd()
	if err != nil {
		return 0, 0, err
	}

	cmd := exec.Command("go", "build", "main.go")
	cmd.Dir = path.Join(base, dir)
	if err := cmd.Run(); err != nil {
		return 0, 0, err
	}

	var totalBytes int64
	for i := 0; i < testRuns; i++ {
		cmd = exec.Command("./main")
		cmd.Dir = path.Join(base, dir)
		buf := &bytes.Buffer{}
		cmd.Stdout = buf
		if err := cmd.Run(); err != nil {
			return 0, 0, err
		}

		nBytes, err := strconv.ParseInt(strings.TrimSpace(buf.String()), 10, 64)
		if err != nil {
			return 0, 0, err
		}
		totalBytes += nBytes
	}

	fi, err := os.Stat(path.Join(dir, "main"))
	if err != nil {
		return 0, 0, err
	}

	return fi.Size(), totalBytes / testRuns, nil
}
