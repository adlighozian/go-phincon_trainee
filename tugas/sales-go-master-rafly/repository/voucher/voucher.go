package voucher

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sales-go/model"
)

type repository struct {}

func NewRepository() *repository {
	return &repository{}
}

type Repositorier interface {
	GetList() (listVoucher []model.Voucher, err error)
	GetVoucherByCode(code string) (voucherData model.Voucher, err error)
	Create(req model.VoucherRequest) (newData model.Voucher, err error)
}

func (repo *repository) getLastID() (lastID int, err error) {
	listVoucher, err := repo.GetList()
	if err != nil {
		return
	}

	if len(listVoucher) == 0 {
		lastID = 0
	} else {
		lastID = len(listVoucher)
	}
	return
}

func (repo *repository) GetList() (listVoucher []model.Voucher, err error) {
	reader, err := os.Open("data/voucher.json")
	if err != nil {
		err = errors.New(fmt.Sprintf("[ERROR] os open voucher json : %s", err.Error()))
		return
	}
	// Create new struct decoder
	decoder := json.NewDecoder(reader)
	// Decode reads the next JSON-encoded value from its, input and stores it in the value pointed to by v.
	decoder.Decode(&listVoucher)

	return
}

func (repo *repository) UpdateJSON(listVoucher []model.Voucher) (err error) {
	writerJson, err := os.Create("data/voucher.json")
	if err != nil {
		err = errors.New(fmt.Sprintf("[ERROR] os create voucher txt : %s", err.Error()))
		return
	}
	encode := json.NewEncoder(writerJson)
	encode.Encode(listVoucher)

	writeTxt, err := os.Create("data/voucher.txt")
	if err != nil {
		err = errors.New(fmt.Sprintf("[ERROR] os create voucher txt : %s", err.Error()))
		return
	}
	encodeToTxt := json.NewEncoder(writeTxt)
	encodeToTxt.Encode(listVoucher)

	return
}

func (repo *repository) GetVoucherByCode(code string) (voucherData model.Voucher, err error) {
	listVoucher, err := repo.GetList()
	if err != nil {
		return
	}

	for _, v := range listVoucher {
		if v.Code == code {
			voucherData = v
		}
	}

	emptyStruct := model.Voucher{}
	if voucherData == emptyStruct {
		err = errors.New("Voucher not found")
		return
	}
	return
}

func (repo *repository) Create(req model.VoucherRequest) (newData model.Voucher, err error) {
	listVoucher, err := repo.GetList()
	if err != nil {
		return
	}

	lastID, err := repo.getLastID()
	if err != nil {
		return
	}

	newData = model.Voucher{
		Id:     lastID + 1,
		Code:   req.Code,
		Persen: req.Persen,
	}
	listVoucher = append(listVoucher, newData)
	
	err = repo.UpdateJSON(listVoucher)
	if err != nil {
		return
	}

	return
}
