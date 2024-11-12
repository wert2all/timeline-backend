package enumvalues

type SettingsType string

const (
	Account SettingsType = "ACCOUNT"
)

// Values provides list valid values for Enum.
func (SettingsType) Values() (kinds []string) {
	for _, s := range []SettingsType{Account} {
		kinds = append(kinds, string(s))
	}
	return
}
