package wav

import (
	"encoding/binary"
	"fmt"
	"io"
)

const (
	RIFF = 0x52494646
	WAVE = 0x57415645
)

type Wav struct {
	ChunkID       uint32 //0x52494646
	Size          uint32 //36 + SubChunk2Size
	Format        uint32 //0x57415645
	Subchunk1ID   uint32 //0x666d7420
	Subchunk1Size uint32
	AudioFormat   AudioFormat
	NumChannels   uint16
	SampleRate    uint32
	ByteRate      uint32
	BlockAlign    uint16
	BitsPerSample uint16
	Subchunk2ID   uint32 //0x64617461
	Subchunk2Size uint32
	Data          []byte //Actual sound data
}

func (w *Wav) ReadFrom(r io.Reader) (n int64, err error) {
	binary.Read(r, binary.BigEndian, &w.ChunkID)
	if w.ChunkID != RIFF {
		return 0, fmt.Errorf("Invalid file format")
	}

	binary.Read(r, binary.LittleEndian, &w.Size)
	binary.Read(r, binary.BigEndian, &w.Format)
	binary.Read(r, binary.BigEndian, &w.Subchunk1ID)
	binary.Read(r, binary.LittleEndian, &w.Subchunk1Size)
	binary.Read(r, binary.LittleEndian, &w.AudioFormat)
	binary.Read(r, binary.LittleEndian, &w.NumChannels)
	binary.Read(r, binary.LittleEndian, &w.SampleRate)
	binary.Read(r, binary.LittleEndian, &w.ByteRate)
	binary.Read(r, binary.LittleEndian, &w.BlockAlign)
	binary.Read(r, binary.LittleEndian, &w.BitsPerSample)
	binary.Read(r, binary.BigEndian, &w.Subchunk2ID)
	binary.Read(r, binary.LittleEndian, &w.Subchunk2Size)
	w.Data = make([]byte, w.Subchunk2Size)
	r.Read(w.Data)
	return
}

func (w *Wav) String() string {
	return fmt.Sprintf(
		"Format=%v,Size=%d,AudioFormat=%d,NumChannels=%d,SampleRate=%d,ByteRate=%d",
		uint32ToString(w.Format),
		w.Subchunk2Size,
		w.AudioFormat,
		w.NumChannels,
		w.SampleRate,
		w.ByteRate)
}

func uint32ToString(u uint32) string {
	var tmp [4]byte
	binary.BigEndian.PutUint32(tmp[:], u)
	return string(tmp[:])
}
