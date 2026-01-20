package app

import (
	"fmt"
	"log/slog"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	increaseDecreaseStepSize int32 = 50
	maxResourceValue         int32 = 500
)

type Int32Entry struct {
	widget.BaseWidget

	Entry          *widget.Entry
	Checkbox       *widget.Check
	increaseButton *widget.Button
	decreaseButton *widget.Button
	canvas         fyne.CanvasObject

	value int32
}

// Create a new entry widget for int32 values with an associated checkbox.
func NewInt32Entry(label string) *Int32Entry {
	entry := widget.NewEntry()
	checkbox := widget.NewCheck(label, nil)
	increaseButton := widget.NewButton("+", nil)
	decreaseButton := widget.NewButton("-", nil)

	e := &Int32Entry{
		Entry:          entry,
		Checkbox:       checkbox,
		increaseButton: increaseButton,
		decreaseButton: decreaseButton,
		canvas:         container.NewVBox(container.NewHBox(checkbox, layout.NewSpacer(), increaseButton, decreaseButton), entry),
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

	e.increaseButton.OnTapped = e.Increase
	e.decreaseButton.OnTapped = e.Decrease

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

// Increase the value by a set amount up to the defined maximum.
// First call will increase it up to the nearest multiple of the step size.
func (e *Int32Entry) Increase() {
	v := e.Get() + increaseDecreaseStepSize
	m := v % increaseDecreaseStepSize
	v -= m
	if v > maxResourceValue {
		v = maxResourceValue
	}
	e.Set(v)
}

// Decrease the value by a set amount up to 0.
// First call will decrease it down to the nearest multiple of the step size.
func (e *Int32Entry) Decrease() {
	v := e.Get()
	m := v % increaseDecreaseStepSize
	if m == 0 {
		v -= increaseDecreaseStepSize
	} else {
		v -= m
	}
	if v < 0 {
		v = 0
	}
	e.Set(v)
}
