package models

import "errors"

var (
	ErrInvalidUsername    = errors.New("недопустимое имя пользователя")
	ErrInvalidEmail       = errors.New("недопустимый email")
	ErrInvalidPassword    = errors.New("недопустимый пароль")
	ErrInvalidUserID      = errors.New("недопустимый ID пользователя")
	ErrInvalidAccountID   = errors.New("недопустимый ID счета")
	ErrInvalidAmount      = errors.New("недопустимая сумма")
	ErrInvalidCurrency    = errors.New("недопустимая валюта")
	ErrInvalidCardNumber  = errors.New("недопустимый номер карты")
	ErrInvalidExpiryDate  = errors.New("недопустимая дата истечения срока")
	ErrInvalidCVV         = errors.New("недопустимый CVV")
	ErrInvalidTerm        = errors.New("недопустимый срок")
	ErrInvalidStatus      = errors.New("недопустимый статус")
	ErrInvalidTransaction = errors.New("недопустимая транзакция")
	ErrInvalidPayment     = errors.New("недопустимый платеж")
	ErrInvalidCredit      = errors.New("недопустимый кредит")
) 