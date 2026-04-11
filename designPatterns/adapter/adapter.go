package adapter

import "fmt"

// Logger is the target interface that the client expects.
type Logger interface {
	Log(message string)
}

// XMLLogger is the adaptee (legacy class) with an incompatible interface.
// It logs messages in XML format.
type XMLLogger struct{}

func (l *XMLLogger) LogXML(message string) {
	fmt.Printf("<log>%s</log>\n", message)
}

// XMLToTextLoggerAdapter is the adapter that makes XMLLogger compatible with Logger.
type XMLToTextLoggerAdapter struct {
	xmlLogger *XMLLogger
}

// NewXMLToTextLoggerAdapter creates a new adapter instance.
func NewXMLToTextLoggerAdapter(xmlLogger *XMLLogger) *XMLToTextLoggerAdapter {
	return &XMLToTextLoggerAdapter{xmlLogger: xmlLogger}
}

// Log implements the Logger interface by adapting the XMLLogger's method.
// It simulates converting XML to plain text (for simplicity, we just strip tags).
func (a *XMLToTextLoggerAdapter) Log(message string) {
	// Call the adaptee's method
	a.xmlLogger.LogXML(message)
	// Simulate adaptation: Print a plain text version
	fmt.Printf("Adapted (plain text): %s\n", message)
}
