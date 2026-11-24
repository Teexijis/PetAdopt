package accountService

import (
	"errors"
	"fmt"
	"net/mail"
	"strings"
	"unicode/utf8"
)

func (s *AccountService) validateInfo(fullName, email, password string, phone *string) error {
	fullName = strings.TrimSpace(fullName)
	if fullName == "" {
		return errors.New("ФИО не должно быть пустым")
	}
	nameLen := utf8.RuneCountInString(fullName) // !! не забыть в других сервисах, чьо в го лен считает байты, а мне нужны символы

	if nameLen < s.minNameLen || nameLen > s.maxNameLen {
		return fmt.Errorf("длина ФИО должна быть от %d до %d символов", s.minNameLen, s.maxNameLen)
	}

	if !s.isValidEmail(email) {
		return errors.New("некорректный email")
	}
	passwordLen := utf8.RuneCountInString(password)

	if passwordLen < s.minPasswordLen {
		return fmt.Errorf("пароль должен содержать минимум %d символов", s.minPasswordLen)
	}

	// тут специально len так как bcrypt ограничивет пароль в байтах
	if len(password) > s.maxPasswordLen {
		return fmt.Errorf("пароль не должен превышать %d байт", s.maxPasswordLen)
	}

	if phone != nil && *phone != "" {
		if !isValidPhone(*phone) {
			return errors.New("некорректный номер телефона")
		}
	}
	return nil
}

func (s *AccountService) isValidEmail(email string) bool {
	if len(email) > s.maxEmailLen {
		return false
	}

	address, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}

	if address.Address != email {
		return false
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}

	if len(parts[1]) == 0 || len(parts[1]) > 253 {
		return false
	}

	return true
}

func isValidPhone(phone string) bool {
	phone = strings.TrimPrefix(phone, "+")

	if len(phone) == 0 || len(phone) > 15 {
		return false
	}

	for _, char := range phone {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
