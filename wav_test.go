package wav_test

import (
	"os"
	"testing"

	"github.com/o5h/wav"
)

func TestExample(t *testing.T) {
	file, _ := os.Open("testdata/sample.wav")
	defer file.Close()
	w := &wav.Wav{}
	w.ReadFrom(file)
	if w.String() != "Format=WAVE,Size=3304656,AudioFormat=1,NumChannels=2,SampleRate=48000,ByteRate=192000" {
		t.Log(w)
		t.Fail()
	}
}
