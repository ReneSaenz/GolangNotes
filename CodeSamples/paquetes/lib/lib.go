package lib

var languages map[string]string

// function to initialize the package
func init() {
	languages = make(map[string]string)

	languages["cs"] = "C#"
	languages["js"] = "JavaScript"
	languages["rb"] = "Ruby"
	languages["go"] = "Golang"
}

func Get(key string) string {
	return languages[key]
}

func Add(key, value string) {
	languages[key] = value
}

func GetAll() map[string]string {
	return languages
}
