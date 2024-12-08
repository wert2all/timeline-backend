package enumvalues

type SettingsType string

const (
	SettingsTypeAccount SettingsType = "ACCOUNT"
)

// Values provides list valid values for Enum.
func (SettingsType) Values() (kinds []string) {
	for _, s := range []SettingsType{SettingsTypeAccount} {
		kinds = append(kinds, string(s))
	}
	return
}
