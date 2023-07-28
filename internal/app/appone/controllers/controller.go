package controllers

import (
	"sync"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/rs/zerolog/log"
)

var Translator ut.Translator
var once sync.Once

func GetTrans() ut.Translator {
	once.Do(func() {
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			zh := zh.New()
			en := en.New()
			uni := ut.New(en, zh, en)

			var ok bool
			Translator, ok = uni.GetTranslator("zh")
			if !ok {
				log.Error().Msg("coun't found translator")
				return
			}

			err := zh_translations.RegisterDefaultTranslations(v, Translator)
			if err != nil {
				log.Error().Msgf("register default translations err:%s", err.Error())
			}
		}
	})

	return Translator
}
