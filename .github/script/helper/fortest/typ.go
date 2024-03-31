package fortest

//go:generate card_json_to_code -type=Color
type Color = string

const (
	ColorRed   Color = "red"
	ColorGreen Color = "green"
)

type ThisStruct struct {
	ColorAlias Color `json:"color_alias"`
}

// colorNameMap BuilderCode for Color name map
var colorNameMap = map[Color]string{
	ColorRed:   "ColorRed",
	ColorGreen: "ColorGreen",
}

// builderCodeColor BuilderCode for Color
func builderCodeColor(val Color) string {
	return colorNameMap[val]
}
