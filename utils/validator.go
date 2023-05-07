package utils

import (
	"errors"
	"miniproject/models"

	validator "github.com/go-playground/validator/v10"
)

func EmailValidation(email string) error {
	v := validator.New()

	err := v.Var(email, "required,email")

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func PostAdminValidation(admin models.Admin) error {
	v := validator.New()

	err := v.Struct(admin)

	if err != nil {
		return errors.New(err.Error())
	}

	validateEmail := EmailValidation(admin.Email)

	if validateEmail != nil {
		return validateEmail
	}

	return nil
}

func PostCustValidation(cust models.Customer) error {
	v := validator.New()

	err := v.Struct(cust)

	if err != nil {
		return errors.New(err.Error())
	}

	validateEmail := EmailValidation(cust.Email)

	if validateEmail != nil {
		return validateEmail
	}

	return nil
}
func PostTransactionValidation(trans models.Transaction) error {
	v := validator.New()

	err := v.Struct(trans)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
func PostTransactionDetailValidation(trans models.TransactionDetail) error {
	v := validator.New()

	err := v.Struct(trans)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
func PostRatingValidation(rating models.TransactionUpdateRating) error {
	v := validator.New()

	err := v.Struct(rating)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
func PostStatusValidation(status models.TransactionUpdateStatus) error {
	v := validator.New()

	err := v.Struct(status)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func PostAreaValidation(area models.Area) error {
	v := validator.New()

	err := v.Struct(area)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
func PostCleanerValidation(cleaner models.Cleaner) error {
	v := validator.New()

	err := v.Struct(cleaner)

	if err != nil {
		return errors.New(err.Error())
	}
	validateEmail := EmailValidation(cleaner.Email)

	if validateEmail != nil {
		return validateEmail
	}

	return nil
}
func PostPaymentValidation(payment models.Payment) error {
	v := validator.New()

	err := v.Struct(payment)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
func PostServiceTypeValidation(st models.ServiceType) error {
	v := validator.New()

	err := v.Struct(st)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func PostStoreValidation(store models.Store) error {
	v := validator.New()

	err := v.Struct(store)

	if err != nil {
		return errors.New(err.Error())
	}
	validateEmail := EmailValidation(store.Email)

	if validateEmail != nil {
		return validateEmail
	}

	return nil
}

func PostTeamValidation(team models.Team) error {
	v := validator.New()

	err := v.Struct(team)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

// func UserActValidator(c echo.Context, id int) error {
// 	user := c.Get("user").(*jwt.Token)
// 	claims := user.Claims.(jwt.MapClaims)
// 	claimsID := claims["id"]

// 	if claimsID == id {
// 		return errors.New("message: unauthorize delete or update")
// 	}
// 	return nil
// }
