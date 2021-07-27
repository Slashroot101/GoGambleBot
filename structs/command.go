package structs

type Command struct {
	name   string
	alias  []string
	action func(phrase string) bool
}
