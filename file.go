package u

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type _file struct{}

func File() *_file {
	return &_file{}
}

// ReadLineForFile 逐行读取文本
func (receiver *_file) ReadLineForFile(filename string, callback func(line string) ) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		if len(line) > 0 {
			callback(strings.TrimSpace(string(line)))
		}
	}
	return nil
}
