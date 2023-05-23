package transaction

// import (
// 	"fmt"
// 	"math/rand"
// 	"net/http"
// 	"sales-go/model"
// 	"sales-go/repository/product"
// 	"sales-go/repository/transaction"
// 	"sales-go/repository/voucher"
// 	"time"
// )

// type handler struct {
// 	repo        transaction.Repositorier
// 	productrepo product.Repositorier
// 	voucherrepo voucher.Repositorier
// }

// func NewHandler(
// 	repositorier transaction.Repositorier,
// 	productRepository product.Repositorier,
// 	voucherRepository voucher.Repositorier,
// ) *handler {
// 	return &handler{
// 		repo:        repositorier,
// 		productrepo: productRepository,
// 		voucherrepo: voucherRepository,
// 	}
// }

// func (handler *handler) GetTransactionByNumber(w http.ResponseWriter, r *http.Request) {
// 	var transactionNumber int
// 	fmt.Println("\nInput transaction number : ")
// 	fmt.Scanln(&transactionNumber)
	
// 	result, err := handler.repo.GetTransactionByNumber(transactionNumber)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	for _, v := range result {
// 		fmt.Println("\nTOKO PHINCON")
// 		fmt.Println("Jl. Arteri Pd. Indah - Jakarta")
// 		fmt.Printf("Transaction Number %d\n", v.Transaction.TransactionNumber)
// 		fmt.Println("--------------------------------------\n")
// 		fmt.Printf("%s\t\t\n", v.Transaction.Name)
// 		fmt.Printf("Rp.%0.2f\t\tx%d\n", v.Price, v.Transaction.Quantity)
// 		fmt.Printf("\nDiscount\t\t%0.2f persen\n", v.Transaction.Discount)
// 		fmt.Printf("Total\t\t\tRp.%0.0f\n", v.Transaction.Total)
// 		fmt.Printf("Pay\t\t\tRp.%0.2f\n", v.Transaction.Pay)
// 		fmt.Printf("Revenue\t\t\tRp.%0.2f\n", v.Transaction.Pay-v.Transaction.Total)
// 	}
// }

// func (handler *handler) CreateBulkTransactionDetail(w http.ResponseWriter, r *http.Request) {	
// 	// 1. input product name
// 	var name string
// 	var quantity int
// 	var total float64
// 	var pay float64
// 	fmt.Println("\nInput product name : ")
// 	fmt.Scanln(&name)

// 	// 2. search product
// 	product, err := handler.productrepo.GetProductByName(name)
// 	if err != nil {	
// 		fmt.Println("\nSorry, the product you are looking for does not exist.")
// 		fmt.Println("Here are the list of products")
// 		handler.productrepo.GetList()

// 		handler.CreateBulkTransactionDetail(w, r)
// 	}

// 	// 3. input quantity
// 	fmt.Println("\nInput quantity : ")
// 	fmt.Scanln(&quantity)

// 	var discount float64
// 	var voucherCode string
// 	var voucherData model.Voucher
// 	if quantity > 300000 {
// 		fmt.Println("\nInput voucher code : ")
// 		fmt.Scanln(&voucherCode)

// 		voucherData, err = handler.voucherrepo.GetVoucherByCode(voucherCode)
// 		if err != nil {
// 			fmt.Println(err)
// 			total = float64(quantity)*product.Price

// 			fmt.Println("Sorry, there is no voucher with name %s\n", voucherCode)
// 		} else {
// 			total = float64(quantity)*product.Price*(voucherData.Persen/100)

// 			fmt.Println("Congratulation, there is a discount : ", voucherData.Persen)
// 		}
// 	} else {
// 		total = float64(quantity)*product.Price
// 	}

// 	// 4. Show total user should pay
// 	fmt.Printf("\nTotal price you should pay : %0.2f\n", total)

// 	// 5. input pay
// 	fmt.Println("\nInput the nominal you want to pay : ")
// 	fmt.Scanln(&pay)
	
// 	// 6. calculate refund
// 	fmt.Println("\nRefund : ", pay - total)

// 	// 7. Input new transaction to transaction detail
// 	newTransaction := model.Transaction{
// 		Id:                len(model.TransactionSlice) + 1,
// 		TransactionNumber: rand.Intn(10000000000),
// 		Name:              name,
// 		Quantity:          quantity,
// 		Discount:          discount,
// 		Total:             total,
// 		Pay:               pay,
// 	}

// 	newTransactionDetail := model.TransactionDetail{
// 		Id:          int(len(model.TransactionSlice)) + 1,
// 		Item:        name,
// 		Price:       product.Price,
// 		Quantity:    quantity,
// 		Total:       total,
// 		Transaction: newTransaction,
// 	}

// 	// 8. Input transaction detail to transaction slice
// 	_, err = handler.repo.CreateBulkTransactionDetail(voucherData, newTransactionDetail, newTransaction)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// 9. Show transaction struct
// 	now := time.Now()
// 	loc, err := time.LoadLocation("Asia/Jakarta")
// 	if err != nil {
// 		panic(err)
// 	}
// 	formatTime := now.In(loc).Format("02/01/2006 15:04")

// 	fmt.Println("\n==== Transaction Success ===\n")
// 	fmt.Println("\nTOKO PHINCON")
// 	fmt.Println("Jl. Arteri Pd. Indah - Jakarta")
// 	fmt.Printf("%s\n", formatTime)
// 	// fmt.Printf("Transaction Number %d\n", result.Transaction.TransactionNumber)
// 	// fmt.Println("--------------------------------------\n")
// 	// fmt.Printf("%s\t\t\n", result.Transaction.Name)
// 	// fmt.Printf("Rp.%0.2f\t\tx%d\n", result.Price, result.Transaction.Quantity)
// 	// fmt.Printf("\nDiscount\t\t%0.2f persen\n", result.Transaction.Discount)
// 	// fmt.Printf("Total\t\t\tRp.%0.0f\n", result.Transaction.Total)
// 	// fmt.Printf("Pay\t\t\tRp.%0.2f\n", result.Transaction.Pay)
// 	// fmt.Printf("Revenue\t\t\tRp.%0.2f\n", result.Transaction.Pay-result.Transaction.Total)
// }