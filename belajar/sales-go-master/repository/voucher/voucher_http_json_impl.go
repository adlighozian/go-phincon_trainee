package voucher

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sales-go/model"
)

type repository struct {}

func NewJsonRepository() *repository {
	return &repository{}
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

func (repo *repository) updateJSON(listVoucher []model.Voucher) (err error) {
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

func (repo *repository) Create(req []model.VoucherRequest) (response []model.Voucher, err error) {
	listVoucher, err := repo.GetList()
	if err != nil {
		return
	}

	for _, v := range req {
		lastID, err := repo.getLastID()
		if err != nil {
			return []model.Voucher{}, err
		}

		newData := model.Voucher{
			Id:     lastID + 1,
			Code:   v.Code,
			Persen: v.Persen,
		}

		listVoucher = append(listVoucher, newData)
		response = listVoucher
		
		err = repo.updateJSON(listVoucher)
		if err != nil {
			return []model.Voucher{}, err
		}
	}

	return
}
