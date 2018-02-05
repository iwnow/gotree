package gotree

// Context is an execution context of GoTree
type context struct {
	Root string
	Args []string
}

// Is need draw files by cmd args
func (c context) isDrawFiles() bool {
	for _, val := range c.Args {
		if val == "-f" {
			return true
		}
	}
	return false
}
