package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	search           string
	fileSearchCount  int
	fileReplaceCount int
)

func main() {
	start := time.Now()
	err := run()
	if err != nil {
		fmt.Println("failed:", err)
	}
	fmt.Printf("finished in %0.2f seconds, modified %d files out of %d scanned that had entries for '%s'\n", time.Since(start).Seconds(), fileReplaceCount, fileSearchCount, search)
}

func run() error {
	var err error

	if len(os.Args) < 2 {
		fmt.Println("usage: dropstrip <search>")
		os.Exit(1)
	}

	search = os.Args[1]

	fmt.Printf("searching for '%s' in *_lde.sql files recursively\n", search)
	err = searchStrip()
	if err != nil {
		return fmt.Errorf("searchStrip: %w", err)
	}

	return nil
}

func searchStrip() error {
	err := filepath.Walk(".", searchStripPath)
	if err != nil {
		return fmt.Errorf("walk: %w", err)
	}
	return nil
}

func searchStripPath(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !strings.HasSuffix(path, "_lde.sql") {
		return nil
	}
	fileSearchCount++
	err = strip(path)
	if err != nil {
		return fmt.Errorf("strip %s: %w", path, err)
	}
	return nil
}

func strip(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	buf := bytes.NewBuffer(nil)
	isModified := false
	for {
		rLine, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		line := string(rLine)
		if strings.Contains(line, search) {
			isModified = true
			continue
		}

		if len(line) == 0 {
			continue
		}
		if !strings.HasPrefix(line, "INSERT") && !strings.HasPrefix(line, "#") {
			continue
		}
		_, err = buf.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("writeString %s: %w", line, err)
		}
	}

	if !isModified {
		return nil
	}
	fileReplaceCount++
	fmt.Println(path)

	f.Close()

	f, err = os.Create(path)
	if err != nil {
		return err
	}

	_, err = f.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}
	return nil
}
