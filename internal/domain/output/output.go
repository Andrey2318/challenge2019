package output

type Model struct {
	DeliveryID string `csv:"delivery id"`
	Status     bool   `csv:"status"`
	Partner    string `csv:"partner"`
	Cost       string `csv:"cost"`
}
