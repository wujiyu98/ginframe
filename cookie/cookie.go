package cookie

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/wujiyu98/ginframe/model"
	"github.com/wujiyu98/ginframe/tool/gaes"
)

var domain string = viper.GetString("servers.domain")
var key string = viper.GetString("servers.key")

func AddCart(ctx *gin.Context, cart model.Cart) error {
	var carts []model.Cart
	cartBase64, err := ctx.Cookie("carts")
	if err == nil {
		cartjson, err := gaes.DecryptString(cartBase64, key)
		if err == nil {
			json.Unmarshal([]byte(cartjson), &carts)
		}
	}
	if err := checkCart(carts, cart); err != nil {
		return err
	}
	carts = append(carts, cart)
	b, _ := json.Marshal(carts)
	value := gaes.EncryptString(string(b), key)
	ctx.SetCookie("carts", value, 9600, "/", domain, false, true)
	return nil
}

func DelCart(ctx *gin.Context, cart model.Cart) error {
	var carts []model.Cart
	cartBase64, err := ctx.Cookie("carts")
	if err == nil {
		cartjson, err := gaes.DecryptString(cartBase64, key)
		if err == nil {
			json.Unmarshal([]byte(cartjson), &carts)
		}
	} else {
		return errors.New("carts is empty")
	}
	carts = delCart(carts, cart)
	if len(carts) > 0 {
		b, _ := json.Marshal(carts)
		value := gaes.EncryptString(string(b), key)
		ctx.SetCookie("carts", value, 9600, "/", domain, false, true)
	} else {
		ctx.SetCookie("carts", "", -1, "/", domain, false, true)
	}

	return nil

}

func delCart(carts []model.Cart, cart model.Cart) []model.Cart {
	var items []model.Cart
	for _, v := range carts {
		if v.ID != cart.ID {
			items = append(items, v)
		}
	}
	return items

}

func checkCart(carts []model.Cart, cart model.Cart) error {
	for _, v := range carts {
		if v.ID == cart.ID {
			return errors.New("cart is exist")
		}
	}
	return nil

}
