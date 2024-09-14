package shipping

type Config struct {
	TPLService     string `env:"TPL_SERVICE" envDefault:"tpl:6565"`
	ManagerService string `env:"MANAGER_SERVICE" envDefault:"manager:6566"`
}
