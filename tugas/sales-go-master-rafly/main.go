package main

import (
	"sales-go/config"
	"fmt"
	"net/http"
	// "sales-go/helpers"

	// handler
	productController "sales-go/handler/product"
	transactionController "sales-go/handler/transaction"
	voucherController "sales-go/handler/voucher"

	// repo
	productRepo "sales-go/repository/product"
	transactionRepo "sales-go/repository/transaction"
	voucherRepo "sales-go/repository/voucher"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err.Error())
	}

	switch config.Storage {
	case "json":
		// repository
		productRepository := productRepo.NewRepository()
		transactionRepository := transactionRepo.NewRepository()
		voucherRepository := voucherRepo.NewRepository()

		// handler
		productHandler := productController.NewHandler(productRepository)
		transactionHandler := transactionController.NewHandler(transactionRepository, productRepository, voucherRepository)
		voucherHandler := voucherController.NewHandler(voucherRepository)

		HTTPServer(config, productHandler, transactionHandler, voucherHandler)
	default:
		// repository
		productRepository := productRepo.NewRepository()
		transactionRepository := transactionRepo.NewRepository()
		voucherRepository := voucherRepo.NewRepository()

		// handler
		productHandler := productController.NewHandler(productRepository)
		transactionHandler := transactionController.NewHandler(transactionRepository, productRepository, voucherRepository)
		voucherHandler := voucherController.NewHandler(voucherRepository)

		Menu(productHandler, transactionHandler, voucherHandler)
	}
}

func HTTPServer(config *config.Config, productHandler productController.Handlerer, transactionHandler transactionController.Handlerer, voucherHandler voucherController.Handlerer) {
	mux := http.NewServeMux()
	mux.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			productHandler.GetList(w, r)
		} else if r.Method == "POST" {
			productHandler.Create(w, r)
		}
	})
	mux.HandleFunc("/voucher", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			voucherHandler.GetList(w, r)
		} else if r.Method == "POST"{
			voucherHandler.Create(w, r)
		}
	})
	mux.HandleFunc("/transaction", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			transactionHandler.GetTransactionByNumber(w, r)
		} else if r.Method == "POST" {
			transactionHandler.CreateBulkTransactionDetail(w, r)
		}
	})

	// create struct server
	server := http.Server{
		Addr: config.Port,
		Handler: mux,
	}

	fmt.Println("Server running on", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	} else if err == nil {
		fmt.Println("Server running on", server.Addr)
	}
}

func Menu(productHandler productController.Handlerer, transactionHandler transactionController.Handlerer, voucherHandler voucherController.Handlerer) {
	var menu int64
	fmt.Println("Choose menu")
	fmt.Println("1. Add New Product")
	fmt.Println("2. Add New Voucher")
	fmt.Println("3. Buy a Product")
	fmt.Println("4. Show List of Product")
	fmt.Println("5. Show List of Voucher")
	fmt.Println("6. Show List of Transaction")
	fmt.Println("7. Get Transaction Detail By Transaction Number")
	fmt.Println("8. Exit")
	fmt.Println("Input menu : ")
	fmt.Scanln(&menu)

	/*
	switch menu {
	case 1:
		productHandler.Create()
		helper.ClearScreeen()
		Menu(productHandler, transactionHandler, voucherHandler)
	case 2:
		voucherHandler.Create()
		helper.ClearScreeen()
		Menu(productHandler, transactionHandler, voucherHandler)
	case 3:
		transactionHandler.CreateTransaction()
		helper.ClearScreeen()
		Menu(productHandler, transactionHandler, voucherHandler)
	case 4:
		productHandler.GetList()
		helper.ClearScreeen()
		Menu(productHandler, transactionHandler, voucherHandler)
	case 5:
		voucherHandler.GetList()
		helper.ClearScreeen()
		Menu(productHandler, transactionHandler, voucherHandler)
	case 6:
		transactionHandler.GetList()
		helper.ClearScreeen()
		Menu(productHandler, transactionHandler, voucherHandler)
	case 7:
		transactionHandler.GetTransactionByNumber()
		helper.ClearScreeen()
		Menu(productHandler, transactionHandler, voucherHandler)
	case 8:
		// Conventionally, code zero indicates success, non-zero an error.
		os.Exit(1)
	}*/
}
