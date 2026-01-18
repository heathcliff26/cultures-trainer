package app

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	fApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/heathcliff26/cultures-trainer/pkg/trainer"
)

const minResourceLabelLength = 25

// Used to change the new app function for testing
var newApp = fApp.New

type App struct {
	app     fyne.App
	main    fyne.Window
	version Version

	startAddress   string
	trainer        *trainer.Trainer
	resourceValues []binding.Int
}

func New() *App {
	app := newApp()
	version := getVersion(app)
	main := app.NewWindow(version.Name)

	a := &App{
		app:            app,
		main:           main,
		version:        version,
		startAddress:   "0x66d0138",
		resourceValues: make([]binding.Int, len(trainer.StorageLocations)),
	}

	a.initContent()

	a.main.SetTitle(version.Name)
	a.main.SetFixedSize(true)
	a.main.Show()
	return a
}

// Simply calls app.Run
func (a *App) Run() {
	a.app.Run()
}

func (a *App) DisplayError(err error) {
	dialog.ShowError(err, a.main)
}

func (a *App) initContent() {
	refreshButton := widget.NewButton("Refresh", a.refreshStorageValues)
	refreshButton.Disable()

	startAddressInput := container.NewVBox(widget.NewLabel("Start Address"), widget.NewEntryWithData(binding.BindString(&a.startAddress)))

	startButton := widget.NewButton("Start Trainer", func() {})
	startButton.OnTapped = func() {
		t, err := trainer.NewTrainer()
		if err != nil {
			a.DisplayError(err)
			return
		}
		a.trainer = t

		startButton.Disable()
		refreshButton.Enable()
		a.refreshStorageValues()
	}

	setupBlock := newBorder(container.NewHBox(startAddressInput, layout.NewSpacer(), startButton))

	storageCategories := newBorder(container.NewHBox(
		a.initStorageCategory("Nahrung", trainer.CategorieNahrung),
		a.initStorageCategory("Bauwaren", trainer.CategorieBauwaren),
		a.initStorageCategory("Resourcen", trainer.CategorieResourcen),
		a.initStorageCategory("Waffen", trainer.CategorieWaffen),
		a.initStorageCategory("Bonusgegenstände", trainer.CategorieBonus),
		a.initStorageCategory("Sonstiges", trainer.CategorieSonstiges),
	))

	a.main.SetContent(container.NewVBox(setupBlock, storageCategories, refreshButton))
}

func (a *App) initStorageCategory(name string, items []string) fyne.CanvasObject {
	obj := make([]fyne.CanvasObject, len(items))
	for i, item := range items {
		index := trainer.StorageIndexes[item]
		a.resourceValues[index] = binding.NewInt()
		labelStr := item
		if len(labelStr) < minResourceLabelLength {
			labelStr += strings.Repeat(" ", minResourceLabelLength-len(labelStr))
		}
		label := widget.NewLabel(labelStr)
		entry := widget.NewEntryWithData(binding.IntToString(a.resourceValues[index]))
		obj[i] = container.NewVBox(label, entry)
	}
	return container.NewVBox(widget.NewLabel(name), newBorder(container.NewVBox(obj...)))
}

func (a *App) refreshStorageValues() {
	if a.trainer == nil {
		return
	}
	hexStr, _ := strings.CutPrefix(a.startAddress, "0x")
	startAddress, err := strconv.ParseUint(hexStr, 16, 64)
	if err != nil {
		a.DisplayError(fmt.Errorf("invalid start address: %v", err))
		return
	}

	values, err := a.trainer.ReadStorageValues(startAddress)
	if err != nil {
		a.DisplayError(fmt.Errorf("failed to read values from game: %v", err))
		return
	}
	for i := range a.resourceValues {
		err = a.resourceValues[i].Set(int(values[i]))
		if err != nil {
			slog.Error("Failed to set resource value", slog.Int("index", i), slog.Int("value", int(values[i])), slog.String("error", err.Error()))
		}
	}
}
