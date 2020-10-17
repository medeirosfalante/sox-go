package sox_test

import (
	"testing"
	"time"

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

func TestTrimJoinFiles(t *testing.T) {
	soxClient := sox.NewSox()
	filepath := "./audios/sample.wav"
	start := float32(2)
	duration := float32(10)
	file1, err := soxClient.Trim(filepath, "./result/output1.wav", start, duration)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if file1 == nil {
		t.Error("file1 is null")
		return
	}
	file2, err := soxClient.Trim(filepath, "./result/output2.wav", float32(5), duration)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if file2 == nil {
		t.Error("file2 is null")
		return
	}
	time.Sleep(2)
	files := []string{file1.FilePath, file2.FilePath}
	joinFiles, errJoin := soxClient.Join(files, "./result/final.wav", true)
	if errJoin != nil {
		t.Errorf("err : %s", errJoin)
		return
	}
	if joinFiles == nil {
		t.Error("joinFiles is null")
		return
	}
}
