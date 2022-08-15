package input

type Model struct {
	DeliveryID string `csv:"Delivery ID"`
	Theatre    string `csv:"Theatre"`
	Size       uint32 `csv:"Size"`
}
