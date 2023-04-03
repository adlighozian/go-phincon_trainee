package handler

type productHandlerHttp struct{}

// func NewProductHandlerHttp() ProductHandlerHttp {
// 	return new(productHandlerHttp)
// }

// http get
// func (repo *productHandlerHttp) ProductGet(write http.ResponseWriter, request *http.Request) {
// 	write.WriteHeader(http.StatusOK)
// 	fmt.Println("Success", http.StatusOK)

// 	contacts := repository.NewContactRepository().List()
// 	result, err := json.Marshal(contacts)

// 	if err != nil {
// 		panic(err)
// 	}
// 	write.WriteHeader(200)
// 	write.Write(result)
// }
