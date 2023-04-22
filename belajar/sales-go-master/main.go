package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/swaggest/swgui/v3emb"

	"sales-go/config"
	"sales-go/docs"
	"sales-go/helpers/middleware"

	// handler
	"sales-go/handler/product"
	"sales-go/handler/transaction"
	"sales-go/handler/voucher"

	// usecase
	productUsecase "sales-go/usecase/product"
	transactionUsecase "sales-go/usecase/transaction"
	voucherUsecase "sales-go/usecase/voucher"

	// repo
	productRepo "sales-go/repository/product"
	transactionRepo "sales-go/repository/transaction"
	voucherRepo "sales-go/repository/voucher"
)

func main() {
	godotenv.Load()

	config, err := config.LoadConfig()
	if err != nil {
		panic(err.Error())
	}

	switch config.Database {
	case "mysql" :
		productRepository := productRepo.NewMySQLHTTPRepository()
		transactionRepository := transactionRepo.NewMySQLHTTPRepository()
		voucherRepository := voucherRepo.NewMySQLHTTPRepository()
		
		productUsecase := productUsecase.NewDBHTTPUsecase(productRepository)
		transactionUsecase := transactionUsecase.NewDBHTTPUsecase(transactionRepository, productRepository, voucherRepository)
		voucherUsecase := voucherUsecase.NewDBHTTPUsecase(voucherRepository)

		productHandler := product.NewDBHTTPHandler(productUsecase)
		transactionHandler := transaction.NewDBHTTPHandler(transactionUsecase)
		voucherHandler := voucher.NewDBHTTPHandler(voucherUsecase)

		DBHTTPServer(config, productHandler, transactionHandler, voucherHandler)
	case "postgresql" :
		productRepository := productRepo.NewPostgreSQLHTTPRepository()
		transactionRepository := transactionRepo.NewPostgreSQLHTTPRepository()
		voucherRepository := voucherRepo.NewPostgreSQLHTTPRepository()
		
		productUsecase := productUsecase.NewDBHTTPUsecase(productRepository)
		transactionUsecase := transactionUsecase.NewDBHTTPUsecase(transactionRepository, productRepository, voucherRepository)
		voucherUsecase := voucherUsecase.NewDBHTTPUsecase(voucherRepository)

		productHandler := product.NewDBHTTPHandler(productUsecase)
		transactionHandler := transaction.NewDBHTTPHandler(transactionUsecase)
		voucherHandler := voucher.NewDBHTTPHandler(voucherUsecase)

		DBHTTPServer(config, productHandler, transactionHandler, voucherHandler)
	case "json":
		productRepository := productRepo.NewJsonRepository()
		transactionRepository := transactionRepo.NewJsonRepository()
		voucherRepository := voucherRepo.NewJsonRepository()

		productHandler := product.NewJsonHTTPHandler(productRepository)
		transactionHandler := transaction.NewJsonHTTPHandler(transactionRepository, productRepository, voucherRepository)
		voucherHandler := voucher.NewJsonHTTPHandler(voucherRepository)

		JsonHTTPServer(config, productHandler, transactionHandler, voucherHandler)
	}
}

func DBHTTPServer(config *config.Config, productHandler product.Handlerer, transactionHandler transaction.Handlerer, voucherHandler voucher.Handlerer) {
	mux := http.NewServeMux()
	mux.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			productHandler.GetList(w, r)
		} else if r.Method == "POST" {
			productHandler.Create(w, r)
		}
	})
	mux.HandleFunc("/voucher", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			voucherHandler.GetList(w, r)
		} else if r.Method == http.MethodPost {
			voucherHandler.Create(w, r)
		}
	})
	mux.HandleFunc("/transaction", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			transactionHandler.GetTransactionByNumber(w, r)
		} else if r.Method == http.MethodPost {
			transactionHandler.CreateBulkTransactionDetail(w, r)
		}
	})

	middleware := middleware.Use(middleware.LoggingHandler(mux))

	// swagger
	docs.SwaggerInfo.Title = "Sales Rest API"
	docs.SwaggerInfo.Description = "Sales Rest API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	mux.Handle("/", v3emb.NewHandler("Sales REST API", "/docs/swagger.json", "/"))

	server := http.Server{
		Addr: config.Port,
		Handler: middleware[0],
	}

	fmt.Println("Server running on", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Test")
	}
}

func JsonHTTPServer(config *config.Config, productHandler product.Handlerer, transactionHandler transaction.Handlerer, voucherHandler voucher.Handlerer) {
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

func Menu(productHandler product.Handlerer, transactionHandler transaction.Handlerer, voucherHandler voucher.Handlerer) {
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
