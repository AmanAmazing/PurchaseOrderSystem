package services

import (
	"PurchaseOrderSystem/models"
	"context"
)

func GetSuppliers() ([]models.Supplier, error) {
	rows, err := db.Query(context.Background(), "SELECT id, name FROM suppliers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var suppliers []models.Supplier
	for rows.Next() {
		var supplier models.Supplier
		err := rows.Scan(&supplier.ID, &supplier.Name)
		if err != nil {
			return nil, err
		}
		suppliers = append(suppliers, supplier)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return suppliers, nil

}
func GetNominals() ([]models.Nominal, error) {
	rows, err := db.Query(context.Background(), "SELECT id, name FROM nominals")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nominals []models.Nominal
	for rows.Next() {
		var nominal models.Nominal
		err := rows.Scan(&nominal.ID, &nominal.Name)
		if err != nil {
			return nil, err
		}
		nominals = append(nominals, nominal)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return nominals, nil
}

func GetProducts() ([]models.Product, error) {
	rows, err := db.Query(context.Background(), "SELECT id, name FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
