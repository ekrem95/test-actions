package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("can't get caller information")
	}

	filepath := path.Join(path.Dir(filename), "time.txt")

	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	s := fmt.Sprintf("%s\n", time.Now().UTC().Format(time.RFC3339))

	if _, err = f.WriteString(s); err != nil {
		panic(err)
	}
}
