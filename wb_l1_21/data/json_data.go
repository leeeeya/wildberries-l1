package data

import "fmt"

// для того, чтобы воспользоваться внешним сервисом для отправки документов в формате XML
// и избежать лишних зависимостей, необходимо применить паттерн "адаптер"

// JsonDocument "внутренний сервис", который работает с данными только json-формата
type JsonDocument struct {
}

// ConvertToXML конвертирует формат документа из json в xml
func (doc JsonDocument) ConvertToXML() string {
	return "<xml></xml>"
}

// JsonDocumentAdapter инициализация адаптера, который наследует JsonDocument
// для избежания связанности внутреннего сервиса с внешним
type JsonDocumentAdapter struct {
	JsonDocument *JsonDocument
}

// SendXMLData реализация интерфейса внешнего сервиса.
// Таким образом, реализация этого метода для JsonDocument будет скрыта от внешнего сервиса
func (adapter JsonDocumentAdapter) SendXMLData() {
	adapter.JsonDocument.ConvertToXML()
	fmt.Println("Sending XML document")
}
