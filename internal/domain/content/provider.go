package content

import (
	"fmt"
)

// GetPage provides text content for requested page in defined language
func GetPage(page, lang string) (map[string]string, error) {
	var content map[string]string
	switch page {
	default:
		return nil, fmt.Errorf("No strings found for (%s) page", page)
	case "home":
		if lang == "ru" {
			content = map[string]string{
				"welcome": "Добро пожаловать в Систему Управления Личными Финансами!",
			}
		} else {
			content = map[string]string{
				"welcome": "Welcome To Personal Finance Management System!",
			}
		}
	}

	return content, nil
}
