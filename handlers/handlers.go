// Package handlers contains the global handlers registry
package handlers

import (
	"bytes"
	"sync"

	"github.com/urfave/cli/v2"
)

var (
	handlers = []*Handler{}

	l = &sync.RWMutex{}
)

type Handler struct {
	Command cli.Command
	Action  HandlerFunc
}

type Context struct {
	*cli.Context
	Input [][]byte
}

func (c *Context) MergeInputs() []byte {
	result := []byte{}

	for _, b := range c.Input {
		result = append(result, b...)
	}

	return result
}

func (c *Context) MergeInputsAsJSON() []byte {
	if len(c.Input) > 1 {
		return append(append([]byte("["), bytes.Join(c.Input, []byte(","))...), []byte("]")...)
	}

	return c.MergeInputs()
}

type HandlerFunc func(*Context) ([]byte, error)

// Handle register a new handler
func Handle(h *Handler) {
	l.Lock()
	defer l.Unlock()

	handlers = append(handlers, h)
}

// Handlers retruns a list with all registered handlers
func Handlers() []*Handler {
	l.RLock()
	defer l.RUnlock()

	result := make([]*Handler, len(handlers))

	copy(result, handlers)

	return result
}
