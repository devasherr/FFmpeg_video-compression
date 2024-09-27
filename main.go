package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func compressVideo(input, output string) error {
	if !fileExists(input) {
		return fmt.Errorf("input file does not exist: %s", input)
	}

	cmd := exec.Command("ffmpeg", "-i", input, "-vcodec", "libx265", "-crf", "28", output)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to compress video: %v, %s", err, stderr.String())
	}

	return nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func main() {
	input := flag.String("input", "", "name of input file")
	output := flag.String("output", "", "name of output file")
	flag.Parse()

	if *input == "" || *output == "" {
		log.Fatal("input and output file is required")
	}

	err := compressVideo(*input, *output)
	if err != nil {
		fmt.Printf("error compressing video: %v", err)
	} else {
		fmt.Println("video compressed successfully")
	}
}
