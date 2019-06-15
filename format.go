package wav

type AudioFormat uint16

func (f AudioFormat) String() string {
	switch f {
	case PCM:
		return "PCM"
	default:
		return "Unknown"
	}
}

const (
	PCM AudioFormat = 1
)
