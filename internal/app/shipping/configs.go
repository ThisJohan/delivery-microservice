package shipping

type Config struct {
	TPLService string `env:"TPL_SERVICE" envDefault:"tpl:6565"`
}
