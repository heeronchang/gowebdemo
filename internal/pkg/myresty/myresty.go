package myresty

import "github.com/go-resty/resty/v2"

func GetClient() *resty.Client {
	client := resty.New()

	return client
}
