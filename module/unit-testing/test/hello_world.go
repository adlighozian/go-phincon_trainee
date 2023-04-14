package helper

func HelloWorld(name string) string {
	return "hello " + name
}

func SapaNama(name string) bool {
	switch {
	case name == "adli":
		return true
	default:
		return false
	}

}
