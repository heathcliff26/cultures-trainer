package app

import (
	"fmt"
	"log/slog"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Int32Entry struct {
	widget.BaseWidget

	Entry    *widget.Entry
	Checkbox *widget.Check
	canvas   fyne.CanvasObject

	value int32
}

// Create a new entry widget for int32 values with an associated checkbox.
func NewInt32Entry(label string) *Int32Entry {
	entry := widget.NewEntry()
	checkbox := widget.NewCheck(minimumLengthString(label, minResourceLabelLength), nil)

	e := &Int32Entry{
		Entry:    entry,
		Checkbox: checkbox,
		canvas:   container.NewVBox(checkbox, entry),
	}

	e.Entry.Validator = func(s string) error {
		_, err := strconv.ParseInt(s, 10, 32)
		return err
	}
	e.Entry.OnChanged = func(s string) {
		if s == "" {
			e.value = 0
			return
		}
		i, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			slog.Error("Failed to parse int32 from entry", slog.String("input", s), slog.String("error", err.Error()))
			return
		}
		e.value = int32(i)
	}
	e.Set(0)

	e.ExtendBaseWidget(e)
	return e
}

func (e *Int32Entry) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(e.canvas)
}

// Set the int32 value of the entry.
func (e *Int32Entry) Set(v int32) {
	e.value = v
	e.Entry.SetText(fmt.Sprintf("%d", e.value))
}

// Get the int32 value of the entry.
func (e *Int32Entry) Get() int32 {
	return e.value
}
