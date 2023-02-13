package envreader

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/mcuadros/go-defaults"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func BindEnv(dest interface{}) (err error) {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigFile(".env")
	v.SetConfigType("env")

	v.AutomaticEnv()

	t := reflect.TypeOf(dest).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := field.Tag.Get("env")
		_ = v.BindEnv(value)
	}

	_ = v.ReadInConfig()

	defaults.SetDefaults(dest)

	err = v.Unmarshal(dest, func(config *mapstructure.DecoderConfig) {
		config.TagName = "env"
	})
	if err != nil {
		return
	}

	validate := validator.New()
	err = validate.Struct(dest)
	return
}
