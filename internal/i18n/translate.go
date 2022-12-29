package i18n

import (
	"context"
)

type Translate interface {
	Trans(ctx context.Context, message string) string
}

type translate struct{}

func New() Translate {
	return &translate{}
}

func (t *translate) Trans(ctx context.Context, message string) string {

	lang := ctx.Value(KeyContext("language"))

	if message == `The "Email" field is not filled` && lang == "ru" {
		return `Поле "Электронная почта" не заполнено`
	} else if message == `The email address is incorrect` && lang == "ru" {
		return "Адрес электронной почты указан неверно"
	} else if message == `The "Password" field is not filled` && lang == "ru" {
		return `Поле "Пароль" не заполнено`
	}

	return message
}
