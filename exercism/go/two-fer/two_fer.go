// Package twofer provides a simple function ShareWith that returns a Two-fer string
// with the name of the person to share with. If the name is empty, the string will
// be "One for you, one for me." Otherwise, the string will be "One for <name>, one for me."
package twofer

// ShareWith return a string with the name of the person to share with.
// If the name is empty, the string will be "One for you, one for me."
// Otherwise, the string will be "One for <name>, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return "One for " + name + ", one for me."
}
