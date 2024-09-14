package manager

type Config struct {
	DeliveryService string `env:"DELIVERY_SERVICE" envDefault:"localhost:6566"`
}
