package database

import(
    "github.com/gin-gonic/gin"
)

var (
    ErrCantFindProduct = errors.New("can't find the product")
    ErrCantDecodeProducts = errors.New("can't find the product")
    ErrUserIdIsNotValid = errors.New("this is not valid user")
    ErrCantUpdateUser = errors.New("cannot add this product to the cart")
    ErrCantRemoveItemCart = errors.New("cannot remove this item from the cart")
    ErrCantGetItem = errors.New("was unable to get the item from the cart")
    ErrCantBuyCartItem  = errors.New("cannot update the purchase")
)

func AddProductToCart(){

}

func RemoveCartItem(){

}

func BuyItemFormcart(){

}

func InstantBuyer(){
    
}