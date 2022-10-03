package data_service

import "fmt"

// AnalyticalDataService "внешний сервис" для работы с данными в формате XML.
// неизменный, всегда отправляет данные в формате XML
type AnalyticalDataService interface {
	SendXMLData()
}

type XMLDocument struct {
}

func (doc XMLDocument) SendXMLData() {
	fmt.Println("Sending XML document")
}
