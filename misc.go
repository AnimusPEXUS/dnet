package dnet

import "regexp"

var VALID_PRESET_NAME_RE *regexp.Regexp = regexp.MustCompile(
	"^[A-Za-z][0-9A-Za-z_]*$",
)

// This is for checking names for Network or Module presets
// return: true - value is valid name
func IsValidPresetName(value string) bool {
	return VALID_PRESET_NAME_RE.Match([]byte(value))
}
