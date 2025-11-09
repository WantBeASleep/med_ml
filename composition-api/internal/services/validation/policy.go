package validation

import (
	"regexp"
	"strings"
)

// ValidatePolicy проверяет корректность номера полиса ОМС
// Поддерживает два формата:
// - Новый формат: 16 цифр
// - Старый формат: серия (3 символа) + номер (6 символов) = 9 символов
func ValidatePolicy(policy string) bool {
	if policy == "" {
		return false
	}

	// Удаляем все пробелы и дефисы
	cleaned := strings.ReplaceAll(strings.ReplaceAll(policy, " ", ""), "-", "")

	// Новый формат: 16 цифр
	newFormat := regexp.MustCompile(`^\d{16}$`)
	if newFormat.MatchString(cleaned) {
		return true
	}

	// Старый формат: 9 символов (3 буквы/цифры + 6 цифр)
	oldFormat := regexp.MustCompile(`^[A-Za-z0-9]{3}\d{6}$`)
	if oldFormat.MatchString(cleaned) {
		return true
	}

	return false
}
