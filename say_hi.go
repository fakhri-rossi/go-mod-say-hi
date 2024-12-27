package go_mod_say_hi

func SayHiToMom() string {
	return "Hi, mom!"
}

// Even the change is just added one parameter, it would break the code
// So it could be determined as major change,
// you should bring this change with new module name,
// as the best practice, you can add the postfix with the new version,
// e.g: github.com/fakhri-rossi/go-mod-say-hi/v2
func SayHi(name string) string {
	return "Hi " + name + "!"
}

func SayHello() string {
	return "Hello!"
}
