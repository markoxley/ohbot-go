package ohbot

type EyeShape struct {
	Name        string
	HexString   string
	AutoMirror  bool
	PupilRangeX float64
	PupilRangeY float64
}

func NewEyeShape(name, hexString string, autoMirror bool, pupilRangeX, pupilRangeY float64) *EyeShape {
	return &EyeShape{
		Name:        name,
		HexString:   hexString,
		AutoMirror:  autoMirror,
		PupilRangeX: pupilRangeX,
		PupilRangeY: pupilRangeY,
	}
}
