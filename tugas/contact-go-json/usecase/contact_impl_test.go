package usecase

import (
	"contact-go/mocks"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

var contactRepository = &mocks.ContactMock{Mock: mock.Mock{}}
var contactUsecase = contactUseCase{Repository: contactRepository}

func TestMain(m *testing.M) {
	fmt.Println("start test")
	start := time.Now()
	m.Run()
	end := time.Since(start)
	fmt.Println("end test")

	fmt.Println("duration", end)
}
