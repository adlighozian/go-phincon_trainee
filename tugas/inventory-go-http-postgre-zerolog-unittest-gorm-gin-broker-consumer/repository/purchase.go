package repository

import (
	"fmt"
	"inventory/db"
	"inventory/model"
)

func TambahPurchase(req []model.SendPurchase) {
	db := db.GetConnection()
	var returns []model.PurchaseDetail
	// Start a transaction
	tx := db.Begin()

	// Execute the native SQL insert statement within the transaction
	insertProduct := `INSERT INTO products (name,price,stock) VALUES ($1,$2,$3)`
	updateProduct := `update products set stock = $1 where name = $2`
	selectProduct := `select stock from products where name = $1`
	insertPurchase := `insert into purchase (order_number,orang,total) values ($1, $2, $3) returning id`
	insertDetail := `insert into purchase_detail (purchase_id, item,price,quantity,total) values ($1,$2,$3,$4,$5) returning id`
	checkProduct := `select id from products where name = $1`

	for _, d := range req {
		var idPurchase uint
		var idPurchaseDetail uint
		var err error
		var idProduct uint
		var stock int
		var resultDetail model.PurchaseDetail

		err = tx.Raw(checkProduct, d.Item).Scan(&idProduct).Error
		if err != nil {
			tx.Rollback()
		}

		if idProduct == 0 {
			fmt.Println("barang belum ada")

			// insert product
			err = tx.Exec(insertProduct, d.Item, d.Price, d.Total).Error
			if err != nil {
				tx.Rollback()
			}

			//insert purchase
			err = tx.Raw(insertPurchase, d.OrderNumber, d.From, d.Total).Scan(&idPurchase).Error
			if err != nil {
				tx.Rollback()
			}

			// insert purchase detail
			err = tx.Raw(insertDetail, idPurchase, d.Item, d.Price, d.Total, d.Total).Scan(&idPurchaseDetail).Error
			if err != nil {
				tx.Rollback()
			}

			resultDetail = model.PurchaseDetail{
				Id:          idPurchaseDetail,
				Purchase_id: idPurchase,
				Item:        d.Item,
				Price:       d.Price,
				Quantity:    d.Total,
				Total:       d.Total,
				Purchase: model.Purchase{
					Id:          idPurchase,
					OrderNumber: d.OrderNumber,
					From:        d.From,
					Total:       d.Total,
				},
			}

			returns = append(returns, resultDetail)

		} else {
			fmt.Println("barang sudah ada")

			// get stock
			err = tx.Raw(selectProduct, d.Item).Scan(&stock).Error
			if err != nil {
				tx.Rollback()
			}
			jumlahStock := stock + d.Total

			// update product
			err = db.Exec(updateProduct, jumlahStock, d.Item).Error
			if err != nil {
				tx.Rollback()
			}

			//insert purchase
			err = tx.Raw(insertPurchase, d.OrderNumber, d.From, d.Total).Scan(&idPurchase).Error
			if err != nil {
				tx.Rollback()
			}

			// insert purchase detail
			err = tx.Raw(insertDetail, idPurchase, d.Item, d.Price, d.Total, jumlahStock).Scan(&idPurchaseDetail).Error
			if err != nil {
				tx.Rollback()
			}

			resultDetail = model.PurchaseDetail{
				Id:          idPurchaseDetail,
				Purchase_id: idPurchase,
				Item:        d.Item,
				Price:       d.Price,
				Quantity:    d.Total,
				Total:       jumlahStock,
				Purchase: model.Purchase{
					Id:          idPurchase,
					OrderNumber: d.OrderNumber,
					From:        d.From,
					Total:       d.Total,
				},
			}

			returns = append(returns, resultDetail)
		}

	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
	}

	// publisher.PubPurchase(returns, "key")
	fmt.Println(returns)
}
