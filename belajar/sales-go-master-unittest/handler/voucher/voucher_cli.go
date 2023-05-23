package voucher

import (
	"fmt"
	"net/http"
	"sales-go/model"
	"sales-go/repository/voucher"
)

type handler struct {
	repo voucher.Repositorier
}

func NewHandler(repositorier voucher.Repositorier) *handler {
	return &handler{
		repo: repositorier,
	}
}

func (handler *handler) GetList(w http.ResponseWriter, r *http.Request) {
	result, err := handler.repo.GetList()
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Printf("\nId\t\t|Code\t\t\t|Persen\t\t")
	for _, v := range result {
		if len(v.Code) > 13 {
			fmt.Printf("\n%d\t\t|%s\t|%0.2f", v.Id, v.Code, v.Persen)
		} else if len(v.Code) > 5 && len(v.Code) < 13 {
			fmt.Printf("\n%d\t\t|%s\t\t|%0.2f", v.Id, v.Code, v.Persen)
		} else {
			fmt.Printf("\n%d\t\t|%s\t\t\t|%0.2f", v.Id, v.Code, v.Persen)
		}
	}
}

func (handler *handler) Create(w http.ResponseWriter, r *http.Request) {	
	var code string
	var persen float64
	fmt.Println("\nInput code data : ")
	fmt.Scanln(&code)
	fmt.Println("\nInput persen data : ")
	fmt.Scanln(&persen)

	
	data := model.VoucherRequest{
		Code:   code,
		Persen: persen,
	}
	result, err := handler.repo.Create([]model.VoucherRequest{data})
	if err != nil {
		fmt.Println(err)
		handler.Create(w, r)
	}

	fmt.Println("Voucher has been added : ", result)
}