# Golang WAV Reader

## Usage 
```go
package main

import (
	"os"

	"github.com/o5h/wav"
)

func main() {
	file, _ := os.Open("testdata/sample.wav")
	defer file.Close()
	w := &wav.Wav{}
	w.ReadFrom(file)
}
```
