package factory

import (
	// _userBusiness "btsid/test/features/users/business"
	// _userData "btsid/test/features/users/data"
	// _userPresentation "pbtsid/testfeatures/users/presentation"

	_authBusiness "btsid/test/features/auth/business"
	_authData "btsid/test/features/auth/data"
	_authPresentation "btsid/test/features/auth/presentation"

	// _productBusiness "project/e-comerce/features/products/bussiness"
	// _productData "project/e-comerce/features/products/data"
	// _productPresentation "project/e-comerce/features/products/presentation"

	// _cartBusiness "project/e-comerce/features/carts/business"
	// _cartData "project/e-comerce/features/carts/data"
	// _cartPresentation "project/e-comerce/features/carts/presentation"
	// _orderBusiness "project/e-comerce/features/orders/bussiness"
	// _orderData "project/e-comerce/features/orders/data"
	// _orderPresentation "project/e-comerce/features/orders/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	// UserPresenter *_userPresentation.UserHandler
	AuthPresenter *_authPresentation.AuthHandler
	// ProductPresenter *_productPresentation.ProductHandler
	// OrderPresenter   *_orderPresentation.OrderHandler
	// CartPresenter    *_cartPresentation.CartHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {

	// userData := _userData.NewUserRepository(dbConn)
	// userBusiness := _userBusiness.NewUserBusiness(userData)
	// userPresentation := _userPresentation.NewUserHandler(userBusiness)

	authData := _authData.NewAuthRepository(dbConn)
	authBusiness := _authBusiness.NewAuthBusiness(authData)
	authPresentation := _authPresentation.NewAuthHandler(authBusiness)

	// productData := _productData.NewProductRepository(dbConn)
	// productBusiness := _productBusiness.NewProductBusiness(productData)
	// productPresentation := _productPresentation.NewProductHandler(productBusiness)

	// cartData := _cartData.NewCartRepository(dbConn)
	// cartBusiness := _cartBusiness.NewCartBusiness(cartData)
	// cartPresentation := _cartPresentation.NewCartHandler(cartBusiness)

	// orderData := _orderData.NewOrderRepository(dbConn)
	// orderBusiness := _orderBusiness.NewOrderBusiness(orderData, cartData)
	// orderPresentation := _orderPresentation.NewOrderHandler(orderBusiness)

	return Presenter{
		//UserPresenter: userPresentation,
		AuthPresenter: authPresentation,
		// ProductPresenter: productPresentation,
		// CartPresenter:    cartPresentation,
		// OrderPresenter:   orderPresentation,
	}
}
