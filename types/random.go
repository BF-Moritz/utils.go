package types

type RandomStringConfig struct {
	Length    uint32
	Lowercase bool
	Uppercase bool
	Numbers   bool
	Symbols   bool
	Override  *string
}
