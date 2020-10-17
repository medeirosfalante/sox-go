package sox

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

// Sox parent group for all functions in sox-go
type Sox struct {
}

// NewSox Create a new instance for sox
func NewSox() Sox {
	return Sox{}
}

//File - return output file define
type File struct {
	FilePath string
}

//Trim - is func for crop the file sound -
func (s Sox) Trim(file string, outputFile string, start float32, duration float32) (*File, error) {
	cmd := exec.Command("sox", file, outputFile, "trim", fmt.Sprintf("%f", start), fmt.Sprintf("%f", duration))
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	log.Printf("in all caps: %q\n", out.String())
	return &File{
		FilePath: outputFile,
	}, nil
}

func (s Sox) Join(files []string, outputFile string, mix bool) (*File, error) {
	commandCobine := "-M"
	if mix {
		commandCobine = "-m"
	}
	params := []string{commandCobine}
	for _, file := range files {
		params = append(params, file)
	}
	params = append(params, outputFile)
	cmd := exec.Command("sox", params...)
	log.Printf("ref %s", cmd.String())
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	log.Printf(" err   %s err", out.String())
	if err != nil {
		return nil, err
	}
	return &File{
		FilePath: outputFile,
	}, nil
}
