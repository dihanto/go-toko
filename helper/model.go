package helper

import (
	"time"

	"github.com/dihanto/go-toko/model/entity"
	"github.com/dihanto/go-toko/model/web/response"
)

func ToResponseCustomerRegister(customer entity.Customer) response.CustomerRegister {
	return response.CustomerRegister{
		Email:        customer.Email,
		Name:         customer.Name,
		RegisteredAt: time.Unix(int64(customer.RegisteredAt), 0),
	}
}

func ToResponseCustomerUpdate(customer entity.Customer) response.CustomerUpdate {
	return response.CustomerUpdate{
		Email:        customer.Email,
		Name:         customer.Name,
		RegisteredAt: time.Unix(int64(customer.RegisteredAt), 0),
		UpdatedAt:    time.Unix(int64(customer.UpdatedAt), 0),
	}
}

func ToResponseSellerRegister(seller entity.Seller) response.SellerRegister {
	return response.SellerRegister{
		Email:        seller.Email,
		Name:         seller.Name,
		RegisteredAt: time.Unix(int64(seller.RegisteredAt), 0),
	}
}

func ToResponseSellerUpdate(seller entity.Seller) response.SellerUpdate {
	return response.SellerUpdate{
		Email:        seller.Email,
		Name:         seller.Name,
		RegisteredAt: time.Unix(int64(seller.RegisteredAt), 0),
		UpdatedAt:    time.Unix(int64(seller.UpdatedAt), 0),
	}
}

func ToResponseAddProduct(product entity.Product) response.AddProduct {
	return response.AddProduct{
		Id:        product.Id,
		IdSeller:  product.IdSeller,
		Name:      product.Name,
		Price:     product.Price,
		Quantity:  product.Quantity,
		CreatedAt: time.Unix(int64(product.CreatedAt), 0),
	}
}

func ToResponseFindById(product entity.Product) response.FindById {
	return response.FindById{
		Id:        product.Id,
		IdSeller:  product.IdSeller,
		Name:      product.Name,
		Price:     product.Price,
		Quantity:  product.Quantity,
		CreatedAt: time.Unix(int64(product.CreatedAt), 0),
		UpdatedAt: time.Unix(int64(product.UpdatedAt), 0),
	}
}

func ToResponseUpdateProduct(product entity.Product) response.UpdateProduct {
	return response.UpdateProduct{
		Id:        product.Id,
		Name:      product.Name,
		Price:     product.Price,
		Quantity:  product.Quantity,
		CreatedAt: time.Unix(int64(product.CreatedAt), 0),
		UpdatedAt: time.Unix(int64(product.UpdatedAt), 0),
	}
}

func ToResponseAddWallet(wallet entity.Wallet) response.AddWallet {
	return response.AddWallet{
		Id:         wallet.Id,
		IdCustomer: wallet.IdCustomer,
		Balance:    wallet.Balance,
		CreatedAt:  time.Unix(int64(wallet.CreatedAt), 0),
	}
}

func ToResponseGetWallet(wallet entity.Wallet) response.GetWallet {
	return response.GetWallet{
		Id:         wallet.Id,
		IdCustomer: wallet.IdCustomer,
		Balance:    wallet.Balance,
		CreatedAt:  time.Unix(int64(wallet.CreatedAt), 0),
		UpdatedAt:  time.Unix(int64(wallet.UpdatedAt), 0),
	}
}

func ToResponseUpdateWallet(wallet entity.Wallet) response.UpdateWallet {
	return response.UpdateWallet{
		IdCustomer: wallet.IdCustomer,
		Balance:    wallet.Balance,
		CreatedAt:  time.Unix(int64(wallet.CreatedAt), 0),
		UpdatedAt:  time.Unix(int64(wallet.UpdatedAt), 0),
	}
}

func ToResponseAddOrder(order entity.Order, orderDetail entity.OrderDetail) response.AddOrder {
	return response.AddOrder{
		Id:         order.Id,
		IdProduct:  orderDetail.IdProduct,
		IdCustomer: order.IdCustomer,
		Quantity:   orderDetail.Quantity,
		OrderedAt:  time.Unix(int64(order.OrderedAt), 0),
	}
}

func ToReponseAddProductToWishlist(product entity.Product) response.AddProductToWishlist {
	return response.AddProductToWishlist{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
		Wishlist: product.Wishlist,
	}
}

func ToReponseDeleteProductFromWishlist(product entity.Product) response.DeleteProductFromWishlist {
	return response.DeleteProductFromWishlist{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
		Wishlist: product.Wishlist,
	}
}
