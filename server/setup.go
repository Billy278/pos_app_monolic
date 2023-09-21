package server

import (
	"fmt"

	db "github.com/Billy278/pos_app_monolic/db"
	ctrlCategories "github.com/Billy278/pos_app_monolic/modules/controllers/categories"
	ctrlPayment "github.com/Billy278/pos_app_monolic/modules/controllers/payment"
	ctrlProducts "github.com/Billy278/pos_app_monolic/modules/controllers/products"
	repoCategories "github.com/Billy278/pos_app_monolic/modules/repository/categories"
	repoPayment "github.com/Billy278/pos_app_monolic/modules/repository/payment"
	repoProducts "github.com/Billy278/pos_app_monolic/modules/repository/products"
	srvCategories "github.com/Billy278/pos_app_monolic/modules/services/categories"
	srvPayment "github.com/Billy278/pos_app_monolic/modules/services/payment"
	srvProducts "github.com/Billy278/pos_app_monolic/modules/services/products"
	"github.com/go-playground/validator/v10"
)

type Handlers struct {
	PaymentCtrl    ctrlPayment.CtrlPayment
	CategoriesCtrl ctrlCategories.CtrlCategories
	ProductsCtrl   ctrlProducts.CtrlProduct
}

func initHandler() Handlers {
	v := validator.New()
	fmt.Println("setup Repository")
	dataStore := db.NewDBPostges()
	repoPayment := repoPayment.NewPaymentImpl(dataStore)
	repoCategory := repoCategories.NewCategoriesImpl(dataStore)
	repoProduct := repoProducts.NewProductsImpl(dataStore)

	fmt.Println("setup services")
	servPayment := srvPayment.NewPaymentSrvImpl(repoPayment)
	servCategory := srvCategories.NewSrvCategoriesImpl(repoCategory)
	servProduct := srvProducts.NewSrvProductImpl(repoProduct, servCategory)

	fmt.Println("setup controllers")
	ctlpayment := ctrlPayment.NewCtrlPaymentImpl(servPayment, v)
	ctlCategory := ctrlCategories.NewCtrlCategoriesimpl(servCategory, v)
	ctlProduct := ctrlProducts.NewCtrlProductimpl(servProduct, v)

	return Handlers{
		PaymentCtrl:    ctlpayment,
		CategoriesCtrl: ctlCategory,
		ProductsCtrl:   ctlProduct,
	}
}
