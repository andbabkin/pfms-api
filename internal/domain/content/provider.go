package content

// GetHomePage provides text content for home page in defined language
func GetHomePage(lang string) (map[string]map[string]string, error) {
	content := make(map[string]map[string]string)
	if lang == "ru" {
		content["home"] = map[string]string{
			"welcome": "Добро пожаловать в Систему Управления Личными Финансами!",
		}
	} else {
		content["home"] = map[string]string{
			"welcome": "Welcome To Personal Finance Management System!",
		}
	}

	return content, nil
}
