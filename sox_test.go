package sox_test

import (
	"testing"

	sox "github.com/rafaeltokyo/sox-go"
)

func TestTrimAudio(t *testing.T) {
	soxClient := sox.NewSox()
	filepath := "./audios/sample.wav"
	filepathOutput := "./result/output.wav"
	start := float32(2)
	duration := float32(10)
	file, err := soxClient.Trim(filepath, filepathOutput, start, duration)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if file == nil {
		t.Error("file is null")
		return
	}
}
