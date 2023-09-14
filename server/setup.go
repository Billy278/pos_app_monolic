package server

import (
	"fmt"

	db "github.com/Billy278/pos_app_monolic/db"
	ctrlPayment "github.com/Billy278/pos_app_monolic/modules/controllers/payment"
	repoPayment "github.com/Billy278/pos_app_monolic/modules/repository/payment"
	srvPayment "github.com/Billy278/pos_app_monolic/modules/services/payment"
	"github.com/go-playground/validator/v10"
)

type Handlers struct {
	PaymentCtrl ctrlPayment.CtrlPayment
}

func initHandler() Handlers {
	v := validator.New()
	fmt.Println("setup Repository")
	dataStore := db.NewDBPostges()
	repoPayment := repoPayment.NewPaymentImpl(dataStore)

	fmt.Println("setup services")
	servPayment := srvPayment.NewPaymentSrvImpl(repoPayment)

	fmt.Println("setup controllers")
	ctlpayment := ctrlPayment.NewCtrlPaymentImpl(servPayment, v)

	return Handlers{
		PaymentCtrl: ctlpayment,
	}
}
