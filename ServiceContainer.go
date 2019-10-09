package main

import (
	"github.com/code-challenge/app/Handlers"
	"github.com/code-challenge/app/controllers"
	"github.com/code-challenge/app/services"
	"sync"
)

type kernel struct{}

func (k *kernel) InjectPaymentController() *controllers.PaymentController {

	paymentService := services.NewPaymentService(Handlers.NewProviderHandler())
	paymentController := controllers.NewPaymentController(paymentService)

	return paymentController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() *kernel {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
