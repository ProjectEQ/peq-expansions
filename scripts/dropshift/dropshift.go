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
	replace          string
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
		fmt.Println("usage: dropshift <pattern>")
		os.Exit(1)
	}

	pattern := strings.Join(os.Args[1:], " ")
	if !strings.HasPrefix(pattern, "(:") {
		return fmt.Errorf("need to start with '(:'")
	}

	if !strings.Contains(pattern, " # ") {
		return fmt.Errorf("need to contain comment")
	}
	replace = pattern[2:]
	replace = replace[0:strings.Index(replace, ":")]

	if strings.Contains(search, "), #") {
		search = strings.ReplaceAll(search, "), #", "); #")
	}

	search = pattern[strings.Index(pattern, ",")+1:]

	fmt.Printf("searching for '%s' in *_lde.sql files recursively and replacing with '%s'\n", search, replace)
	err = searchShift()
	if err != nil {
		return fmt.Errorf("searchStrip: %w", err)
	}

	return nil
}

func searchShift() error {
	err := filepath.Walk(".", searchShiftPath)
	if err != nil {
		return fmt.Errorf("walk: %w", err)
	}
	return nil
}

func searchShiftPath(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !strings.HasSuffix(path, "_lde.sql") {
		return nil
	}
	fileSearchCount++
	err = shift(path)
	if err != nil {
		return fmt.Errorf("shift %s: %w", path, err)
	}
	return nil
}

func shift(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w, err := os.OpenFile(strings.ReplaceAll(path, "_lde.sql", "_lte.sql"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer w.Close()

	zone := ""

	zone = path[0:strings.Index(path, "_")]
	if strings.Contains(zone, "/") {
		zone = zone[strings.Index(zone, "/")+1:]
	}

	npcsOut := make(map[string]bool)
	r := bufio.NewReader(f)
	buf := bytes.NewBuffer(nil)
	isModified := false
	lineNumber := 0
	for {
		lineNumber++
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
			//INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES(:todo:, 7161, 100); # Partially Digested Root 12436 20396 8.75 a_xakra_larva (155015 lvl 3 chance 8.75%), a_xakra_larva (155027 lvl 1 chance 8.75%), a_xakra_larva (155028 lvl 2 chance 8.75%)
			shift := line
			pattern := "; # "
			pos := strings.Index(shift, pattern)
			if pos < 1 {
				return fmt.Errorf("couldn't find pattern '%s'", pattern)
			}
			shift = shift[pos+3:]

			npcs := strings.Split(shift, "chance")
			for i := range npcs {
				npc := npcs[i]
				if !strings.Contains(npc, "lvl") {
					continue
				}
				npc = reverse(npc)
				pattern = " "

				pos = strings.Index(npc, pattern)
				if pos < 0 {
					return fmt.Errorf("couldn't find space in npc 1 %d '%s'", lineNumber, npc)
				}

				npc = npc[pos+1:]
				pos = strings.Index(npc, pattern)
				if pos < 0 {
					return fmt.Errorf("couldn't find space in npc 2 %d '%s'", lineNumber, npc)
				}
				npc = npc[pos+1:]
				pos = strings.Index(npc, pattern)
				if pos < 0 {
					return fmt.Errorf("couldn't find space in npc 3 %d '%s'", lineNumber, npc)
				}
				npc = npc[pos+1:]
				pos = strings.Index(npc, pattern)
				if pos < 0 {
					return fmt.Errorf("couldn't find space in npc 4 %d '%s'", lineNumber, npc)
				}
				npc = npc[pos+1:]

				pos = strings.Index(npc, pattern)
				if pos < 0 {
					return fmt.Errorf("couldn't find space in npc 4 %d '%s'", lineNumber, npc)
				}
				npc = npc[0:pos]
				npc = strings.TrimSpace(reverse(npc))
				npc = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(npc), " ", "_"), "#", ""), "`", "")
				npcsOut[npc] = true
			}
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

	for npc := range npcsOut {
		_, err := w.WriteString(fmt.Sprintf("(:%s_%s_lt:, :%s:, 1),\n", zone, npc, replace))
		if err != nil {
			return fmt.Errorf("inject npc %s: %s", npc, err)
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

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
