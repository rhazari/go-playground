package adapter

import "testing"

func Test_Adapter(t *testing.T) {
	// Create the adaptee (legacy XML logger)
	xmlLogger := &XMLLogger{}

	// Create the adapter to make it compatible with Logger
	adapter := NewXMLToTextLoggerAdapter(xmlLogger)

	// Now use the adapter as if it were a Logger
	adapter.Log("This is a test message")
	// Output:
	// <log>This is a test message</log>
	// Adapted (plain text): This is a test message
}