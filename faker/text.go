package faker

import "strings"

var words = []string{
	"hawk",
	"go",
	"framework",
	"developer",
	"backend",
	"fast",
	"simple",
	"clean",
	"modern",
	"secure",
}

func Word() string {
	return random(words)
}

func Sentence() string {

	var result []string

	length := Int(5, 12)

	for i := 0; i < length; i++ {
		result = append(result, Word())
	}

	return strings.Join(result, " ") + "."
}

func Paragraph() string {

	count := Int(3, 7)

	result := ""

	for i := 0; i < count; i++ {

		result += Sentence()

		if i != count-1 {
			result += " "
		}
	}

	return result
}

func Slug() string {

	return strings.ToLower(
		strings.ReplaceAll(
			Sentence(),
			" ",
			"-",
		),
	)
}

func Company() string {

	companies := []string{
		"OpenAI",
		"Hawk",
		"TechNova",
		"Byte Labs",
		"CloudStack",
		"DevCore",
		"Acme",
	}

	return random(companies)
}

func URL() string {

	return "https://www." + Username() + ".com"
}

func Title() string {

	titles := []string{
		"Software Engineer",
		"Backend Developer",
		"Frontend Developer",
		"Product Manager",
		"DevOps Engineer",
	}

	return random(titles)
}