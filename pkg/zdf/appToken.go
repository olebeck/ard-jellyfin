package zdf

import (
	"fmt"
	"io"
	"regexp"
)

var appTokenRegex = regexp.MustCompile(`\\"appToken\\":{\\"apiToken\\":\\"([^"]*)\\"`)

func GetAppToken() (string, error) {
	res, err := client.Get("https://www.zdf.de/live-tv")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	match := appTokenRegex.FindSubmatch(body)
	if match == nil {
		return "", fmt.Errorf("appToken not found")
	}

	return string(match[1]), nil
}
