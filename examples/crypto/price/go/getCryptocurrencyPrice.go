package example

import (
	"fmt"
	"github.com/micro/services/clients/go/crypto"
	"os"
)

// Get the last price for a given crypto ticker
func GetCryptocurrencyPrice() {
	cryptoService := crypto.NewCryptoService(os.Getenv("MICRO_API_TOKEN"))
	rsp, err := cryptoService.Price(&crypto.PriceRequest{
		Symbol: "BTCUSD",
	})
	fmt.Println(rsp, err)
}