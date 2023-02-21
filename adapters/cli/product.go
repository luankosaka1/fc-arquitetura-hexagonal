package cli

import (
	"fmt"
	"github.com/luankosaka1/arquitetura-hexagonal-golang/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("product ID %s with the name %s has been created", product.GetId(), product.GetName())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("product %s has been enabled", res.GetName())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("product %s has been disabled", res.GetName())
	default:
		res, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("product #%s with name %s", res.GetId(), res.GetName())
	}

	return result, nil
}
