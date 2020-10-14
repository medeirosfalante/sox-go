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
		FilePath: out.String(),
	}, nil
}
