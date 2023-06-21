package utils

import (
	"os"
	"strings"
	"time"
)

func DoWithTries(fn func() error, attempts int, duration time.Duration) (err error) {
	for attempts > 0 {
		err = fn()
		if err != nil {
			time.Sleep(duration)
			attempts--
			continue
		}
		return nil
	}
	return err
}

func FormatQuery(query string) string {
	return strings.ReplaceAll(strings.ReplaceAll(query, "\t", " "), "\n", " ")
}

// func TranslitURL(url string) string {
// 	var baseRuEn = map[string]string{
// 		"а": "a", "А": "A", "Б": "B", "б": "b", "В": "V", "в": "v", "Г": "G", "г": "g",
// 		"Д": "D", "д": "d", "З": "Z", "з": "z", "И": "I", "и": "i", "К": "K", "к": "k",
// 		"Л": "L", "л": "l", "М": "M", "м": "m", "Н": "N", "н": "n", "О": "O", "о": "o",
// 		"П": "P", "п": "p", "Р": "R", "р": "r", "С": "S", "с": "s", "Т": "T", "т": "t",
// 		"У": "U", "у": "u", "Ф": "F", "ф": "f",
// 	}
// 	var result string
// 	urlByte := []byte(url)
// 	for _, v := range urlByte {
// 		fmt.Println(string(v))
// 		n := baseRuEn[string(v)]
// 		fmt.Println(n)
// 		result += n
// 	}
// 	fmt.Println(result)
// 	return result
// }

func TrimSpace(url string) string {
	return strings.ReplaceAll(url, " ", "_")
}

func DownloadBook(pathPDF string, userFile string) error {

	r, err := os.Open(pathPDF)
	if err != nil {
		return err
	}
	defer r.Close()
	w, err := os.Create(userFile)
	if err != nil {
		return err
	}
	defer w.Close()
	_, err = w.ReadFrom(r)
	if err != nil {
		return err
	}

	return nil
}
