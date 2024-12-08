package cli

import (
	"fmt"

	"github.com/lucasantoniooficial/hexagonal-go/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		product, err = service.Enable(product)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID %s has been enabled", product.GetID())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		product, err = service.Disable(product)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID %s has been disabled", product.GetID())
	default:
		products, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID %s with the name %s has the price %f and status %s", products.GetID(), products.GetName(), products.GetPrice(), products.GetStatus())
	}

	return result, nil
}
