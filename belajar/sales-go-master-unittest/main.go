package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	// "github.com/swaggest/swgui/v3emb"

	"sales-go/db"
	"sales-go/config"
	"sales-go/helpers/clearscreen"
	"sales-go/helpers/middleware"
	"sales-go/helpers/random"
	"sales-go/publisher"

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

	switch config.App {
	case "postgresql-gin-gorm":		
		db, err := client.NewConnection(client.Database).GetDBGormConnection()
		if err != nil {
			panic(err)
		}

		DB, err := db.DB()
		if err != nil {
			panic(err)
		}

		publisher := publisher.NewPublisher()
		random := random.NewRandom()

		productRepository := productRepo.NewPostgreSQLHTTPRepository(DB)
		transactionRepository := transactionRepo.NewPostgreSQLHTTPRepository(DB, publisher, random)
		voucherRepository := voucherRepo.NewPostgreSQLHTTPRepository(DB)

		productUsecase := productUsecase.NewDBHTTPUsecase(productRepository)
		transactionUsecase := transactionUsecase.NewDBHTTPUsecase(transactionRepository, productRepository, voucherRepository)
		voucherUsecase := voucherUsecase.NewDBHTTPUsecase(voucherRepository)

		productHandler := product.NewGinDBHTTPHandler(productUsecase)
		transactionHandler := transaction.NewGinDBHTTPHandler(transactionUsecase)
		voucherHandler := voucher.NewGinDBHTTPHandler(voucherUsecase)

		DBGinHTTPServer(config, productHandler, transactionHandler, voucherHandler)
	case "postgresql-gin":		
		db, err := client.NewConnection(client.Database).GetDBConnection()
		if err != nil {
			panic(err)
		}

		publisher := publisher.NewPublisher()
		random := random.NewRandom()

		productRepository := productRepo.NewPostgreSQLHTTPRepository(db)
		transactionRepository := transactionRepo.NewPostgreSQLHTTPRepository(db, publisher, random)
		voucherRepository := voucherRepo.NewPostgreSQLHTTPRepository(db)

		productUsecase := productUsecase.NewDBHTTPUsecase(productRepository)
		transactionUsecase := transactionUsecase.NewDBHTTPUsecase(transactionRepository, productRepository, voucherRepository)
		voucherUsecase := voucherUsecase.NewDBHTTPUsecase(voucherRepository)

		productHandler := product.NewGinDBHTTPHandler(productUsecase)
		transactionHandler := transaction.NewGinDBHTTPHandler(transactionUsecase)
		voucherHandler := voucher.NewGinDBHTTPHandler(voucherUsecase)

		DBGinHTTPServer(config, productHandler, transactionHandler, voucherHandler)
	case "mysql" :
		db, err := client.NewConnection(client.Database).GetDBConnection()
		if err != nil {
			panic(err)
		}

		random := random.NewRandom()

		productRepository := productRepo.NewMySQLHTTPRepository(db)
		transactionRepository := transactionRepo.NewMySQLHTTPRepository(db, random)
		voucherRepository := voucherRepo.NewMySQLHTTPRepository(db)
		
		productUsecase := productUsecase.NewDBHTTPUsecase(productRepository)
		transactionUsecase := transactionUsecase.NewDBHTTPUsecase(transactionRepository, productRepository, voucherRepository)
		voucherUsecase := voucherUsecase.NewDBHTTPUsecase(voucherRepository)

		productHandler := product.NewDBHTTPHandler(productUsecase)
		transactionHandler := transaction.NewDBHTTPHandler(transactionUsecase)
		voucherHandler := voucher.NewDBHTTPHandler(voucherUsecase)

		DBHTTPServer(config, productHandler, transactionHandler, voucherHandler)
	case "postgresql" :
		db, err := client.NewConnection(client.Database).GetDBConnection()
		if err != nil {
			panic(err)
		}
		
		publisher := publisher.NewPublisher()
		random := random.NewRandom()

		productRepository := productRepo.NewPostgreSQLHTTPRepository(db)
		transactionRepository := transactionRepo.NewPostgreSQLHTTPRepository(db, publisher, random)
		voucherRepository := voucherRepo.NewPostgreSQLHTTPRepository(db)
		
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
	case "cli":
		productRepository := productRepo.NewCLIRepository()
		transactionRepository := transactionRepo.NewCLIRepository()
		voucherRepository := voucherRepo.NewCLIRepository()

		productHandler := product.NewJsonHTTPHandler(productRepository)
		transactionHandler := transaction.NewJsonHTTPHandler(transactionRepository, productRepository, voucherRepository)
		voucherHandler := voucher.NewJsonHTTPHandler(voucherRepository)

		Menu(productHandler, transactionHandler, voucherHandler)
	}
}

func DBGinHTTPServer(config *config.Config,productHandler product.GinHandlerer, transactionHandler transaction.GinHandlerer, voucherHandler voucher.GinHandlerer) {
	r := gin.Default()
	r.Use(middleware.HeaderVerificationMiddleware)

	r1 := r.Group("/product")
	r1.GET("/", productHandler.GetList)
	r1.POST("/", productHandler.Create)

	r2 := r.Group("/transaction")
	r2.GET("/", transactionHandler.GetTransactionByNumber)
	r2.POST("/", transactionHandler.CreateBulkTransactionDetail)

	r3 := r.Group("/voucher")
	r3.GET("/", voucherHandler.GetList)
	r3.POST("/", voucherHandler.Create)

	r.Run(config.Port)
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

	middleware := middleware.Use(middleware.LoggingMiddleware(mux))

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

	w := new(http.ResponseWriter)
	r := new(*http.Request)
	switch menu {
	case 1:
		productHandler.Create(*w, *r)
		clearscreen.ClearScreeen()
		Menu(productHandler, transactionHandler, voucherHandler)
	case 2:
		voucherHandler.Create(*w, *r)
		clearscreen.ClearScreeen()
		Menu(productHandler, transactionHandler, voucherHandler)
	case 3:
		transactionHandler.CreateBulkTransactionDetail(*w, *r)
		clearscreen.ClearScreeen()
		Menu(productHandler, transactionHandler, voucherHandler)
	case 4:
		productHandler.GetList(*w, *r)
		clearscreen.ClearScreeen()
		Menu(productHandler, transactionHandler, voucherHandler)
	case 5:
		voucherHandler.GetList(*w, *r)
		clearscreen.ClearScreeen()
		Menu(productHandler, transactionHandler, voucherHandler)
	case 6:
		transactionHandler.GetTransactionByNumber(*w, *r)
		clearscreen.ClearScreeen()
		Menu(productHandler, transactionHandler, voucherHandler)
	case 7:
		// Conventionally, code zero indicates success, non-zero an error.
		os.Exit(1)
	}
}
