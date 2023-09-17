package server

import (
	"fmt"

	db "github.com/Billy278/pos_app_monolic/db"
	ctrlCategories "github.com/Billy278/pos_app_monolic/modules/controllers/categories"
	ctrlPayment "github.com/Billy278/pos_app_monolic/modules/controllers/payment"
	repoCategories "github.com/Billy278/pos_app_monolic/modules/repository/categories"
	repoPayment "github.com/Billy278/pos_app_monolic/modules/repository/payment"
	srvCategories "github.com/Billy278/pos_app_monolic/modules/services/categories"
	srvPayment "github.com/Billy278/pos_app_monolic/modules/services/payment"
	"github.com/go-playground/validator/v10"
)

type Handlers struct {
	PaymentCtrl    ctrlPayment.CtrlPayment
	CategoriesCtrl ctrlCategories.CtrlCategories
}

func initHandler() Handlers {
	v := validator.New()
	fmt.Println("setup Repository")
	dataStore := db.NewDBPostges()
	repoPayment := repoPayment.NewPaymentImpl(dataStore)
	repoCategory := repoCategories.NewCategoriesImpl(dataStore)

	fmt.Println("setup services")
	servPayment := srvPayment.NewPaymentSrvImpl(repoPayment)
	servCategory := srvCategories.NewSrvCategoriesImpl(repoCategory)

	fmt.Println("setup controllers")
	ctlpayment := ctrlPayment.NewCtrlPaymentImpl(servPayment, v)
	ctlCategory := ctrlCategories.NewCtrlCategoriesimpl(servCategory, v)

	return Handlers{
		PaymentCtrl:    ctlpayment,
		CategoriesCtrl: ctlCategory,
	}
}
