package pattern

import "fmt"

// Handler interface
type Handler interface {
	ProcessRequest(fileType string) string
}

// Conrecete handlers
type ImageHandler struct {
	Next Handler
}

func (i *ImageHandler) ProcessRequest(fileType string) string {
	const handlerType = "image"
	if fileType == handlerType {
		return fmt.Sprintf("SUCCES: %s file is processed", fileType)
	} else if i.Next != nil {
		return i.Next.ProcessRequest(fileType)
	}
	return fmt.Sprintf("FAILED: %s file is not processed", fileType)
}

type TextHandler struct {
	Next Handler
}

func (t *TextHandler) ProcessRequest(fileType string) string {
	const handlerType = "text"
	if fileType == handlerType {
		return fmt.Sprintf("SUCCES: %s file is processed", fileType)
	} else if t.Next != nil {
		return t.Next.ProcessRequest(fileType)
	}
	return fmt.Sprintf("FAILED: %s file is not processed", fileType)
}
