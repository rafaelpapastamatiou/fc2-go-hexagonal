package cli

import (
	"fmt"

	"github.com/rafaelpapastamatiou/fc2-go-hexagonal/application"
)

const (
	RESULT_CREATE_MESSAGE  = "Product ID %s with the name %s has been created with price %f and status %s"
	RESULT_ENABLE_MESSAGE  = "Product %s has been enabled."
	RESULT_DISABLE_MESSAGE = "Product %s has been disabled."
	RESULT_GET_MESSAGE     = "Product ID: %s\nName: %s\nPrice: %f\nStatus: %s"
)

func Run(
	service application.ProductServiceInterface,
	action string,
	id string,
	name string,
	price float64,
) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(name, price)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf(
			RESULT_CREATE_MESSAGE,
			product.GetID(),
			product.GetName(),
			product.GetPrice(),
			product.GetStatus(),
		)

	case "enable":
		product, err := service.Get(id)

		if err != nil {
			return result, err
		}

		res, err := service.Enable(product)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf(RESULT_ENABLE_MESSAGE, res.GetName())

	case "disable":
		product, err := service.Get(id)

		if err != nil {
			return result, err
		}

		res, err := service.Disable(product)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf(RESULT_DISABLE_MESSAGE, res.GetName())

	default:
		res, err := service.Get(id)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf(
			RESULT_GET_MESSAGE,
			res.GetID(),
			res.GetName(),
			res.GetPrice(),
			res.GetStatus(),
		)
	}

	return result, nil
}
