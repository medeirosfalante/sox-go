package sox_test

import (
	"os"
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
	os.Remove(filepathOutput)
}

func TestInfo(t *testing.T) {
	soxClient := sox.NewSox()
	filepath := "./audios/sample.wav"
	file, err := soxClient.Info(filepath)
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
	os.Mkdir("./result", 777)
	soxClient := sox.NewSox()
	filepath := "./audios/sample.wav"
	start := float32(2)
	duration := float32(10)
	finalFilePath := "./result/final.wav"
	files := []string{"./result/output1.wav", "./result/output2.wav"}
	file1, err := soxClient.Trim(filepath, files[0], start, duration)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if file1 == nil {
		t.Error("file1 is null")
		return
	}
	file2, err := soxClient.Trim(filepath, files[1], float32(5), duration)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if file2 == nil {
		t.Error("file2 is null")
		return
	}
	time.Sleep(2)
	joinFiles, errJoin := soxClient.Join(files, finalFilePath, true)
	if errJoin != nil {
		t.Errorf("err : %s", errJoin)
		return
	}
	if joinFiles == nil {
		t.Error("joinFiles is null")
		return
	}

	removeFiles := append(files, finalFilePath)

	for _, item := range removeFiles {
		os.Remove(item)
	}

}
