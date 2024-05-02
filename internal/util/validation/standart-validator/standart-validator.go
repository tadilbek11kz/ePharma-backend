package standard_validator

import (
	"regexp"
	"strconv"
	"time"

	"github.com/tadilbek11kz/ePharma-backend/internal/util/validation"

	"github.com/go-playground/validator/v10"
	"github.com/relvacode/iso8601"
)

type StandardValidator struct {
	Validate *validator.Validate
}

func New() validation.Validator {
	standardValidator := &StandardValidator{validator.New()}

	var validators = map[string]validator.Func{
		"is-iin":            iinValidate,
		"is-iso-date":       isISODate,
		"is-not-expired":    isNotExpired,
		"is-tt-code":        isTTCode,
		"has-lang":          hasLang,
		"is-phone":          isPhone,
		"is-postLink-valid": validatePostLink,
		"is-broker-phone":   isBrokerPhone,
	}

	for name, fn := range validators {
		err := standardValidator.Validate.RegisterValidation(name, fn)
		if err != nil {
			return nil
		}
	}

	return standardValidator
}

func (standardValidator *StandardValidator) ValidateStruct(model interface{}) error {
	err := standardValidator.Validate.Struct(model)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Validate IIN
func iinValidate(fl validator.FieldLevel) bool {
	const iinLen = 12

	return len(fl.Field().String()) == iinLen
}

// Check string is ISO8601
func isISODate(fl validator.FieldLevel) bool {
	_, err := iso8601.ParseString(fl.Field().String())

	return err == nil
}

func isNotExpired(fl validator.FieldLevel) bool {
	t, err := iso8601.ParseString(fl.Field().String())
	if err != nil {
		return false
	}

	const bitSize = 64
	param := fl.Param()
	pi, err := strconv.ParseFloat(param, bitSize)
	if err != nil {
		return false
	}

	r := t.After(time.Now().Add(time.Minute * time.Duration(pi)))

	return r
}

// Validate if PostLink or PostLinkCamel is set
func validatePostLink(fl validator.FieldLevel) bool {
	return len(fl.Parent().FieldByName("PostLinkCamel").String()) > 0 || len(fl.Parent().FieldByName("PostLink").String()) > 0
}

// Validate Partner TT StatusCode **-****-**
func isTTCode(fl validator.FieldLevel) bool {
	TTCodeRgxStr := "^..-....-..$"
	TTCodeRgx := regexp.MustCompile(TTCodeRgxStr)

	return TTCodeRgx.MatchString(fl.Field().String())
}

// hasLang validates language field existence
func hasLang(fl validator.FieldLevel) bool {
	langs := []string{"ru", "en", "kk"}

	for _, l := range langs {
		if l == fl.Field().String() {
			return true
		}
	}

	return false
}

// isPhone validates phone number
func isPhone(fl validator.FieldLevel) bool {
	phoneRegexStr := `^7\d{9}$`
	phoneRegex := regexp.MustCompile(phoneRegexStr)

	return phoneRegex.MatchString(fl.Field().String())
}

// isBrokerPhone validates phone number
func isBrokerPhone(fl validator.FieldLevel) bool {
	phoneRegexStr := "^77\\d{9}$"
	phoneRegex := regexp.MustCompile(phoneRegexStr)

	return phoneRegex.MatchString(fl.Field().String())
}
