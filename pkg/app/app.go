package app

import (
	"fmt"
	"log/slog"

	"fyne.io/fyne/v2"
	fApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/heathcliff26/cultures-trainer/pkg/trainer"
)

const minResourceLabelLength = 25

const (
	freezeButtonTextFreeze   = "Freeze"
	freezeButtonTextUnfreeze = "Unfreeze"
)

// Used to change the new app function for testing
var newApp = fApp.New

type App struct {
	app     fyne.App
	main    fyne.Window
	version Version

	trainer *trainer.Trainer

	storageAddressOffset uint64
	storageAddressEntry  *widget.Entry
	resourceValues       []binding.Int
	resourceFreezeChecks []*widget.Check
	freezeButton         *widget.Button
}

func New() *App {
	app := newApp()
	version := getVersion(app)
	main := app.NewWindow(version.Name)

	a := &App{
		app:                  app,
		main:                 main,
		version:              version,
		resourceValues:       make([]binding.Int, len(trainer.StorageLocations)),
		resourceFreezeChecks: make([]*widget.Check, len(trainer.StorageLocations)),
	}

	a.initContent()

	a.main.SetTitle(version.Name)
	a.main.SetFixedSize(true)
	a.main.Resize(fyne.NewSquareSize(400))
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
	a.storageAddressEntry = widget.NewEntry()
	a.storageAddressEntry.SetText("0x0000000")
	a.storageAddressEntry.Validator = func(s string) error {
		_, err := hexStringToUint64(s)
		return err
	}
	storageAddressInput := container.NewVBox(widget.NewLabel(minimumLengthString("Storage Address", 20)), a.storageAddressEntry)

	storageAddressTypeSelect := widget.NewSelect(trainer.StorageLocations, func(s string) {
		offset := trainer.StorageIndexes[s]
		a.storageAddressOffset = uint64(offset)
	})
	storageAddressTypeSelect.SetSelectedIndex(7)
	storageAddressType := container.NewVBox(widget.NewLabel(minimumLengthString("Resource Name", 30)), storageAddressTypeSelect)

	startButton := widget.NewButton("Start Trainer", func() {})

	setupBlock := container.NewHBox(storageAddressType, storageAddressInput)
	setupBlock = container.NewVBox(setupBlock, startButton)

	storageCategories := newBorder(container.NewHBox(
		a.initStorageCategory("Nahrung", trainer.CategorieNahrung),
		a.initStorageCategory("Bauwaren", trainer.CategorieBauwaren),
		a.initStorageCategory("Resourcen", trainer.CategorieResourcen),
		a.initStorageCategory("Waffen", trainer.CategorieWaffen),
		a.initStorageCategory("Bonusgegenstände", trainer.CategorieBonus),
		a.initStorageCategory("Sonstiges", trainer.CategorieSonstiges),
	))

	refreshButton := widget.NewButton("Refresh", a.refreshStorageValues)
	applyButton := widget.NewButton("Apply", a.applyStorageValues)
	a.freezeButton = widget.NewButton(freezeButtonTextFreeze, a.freezeSelectedValues)

	runBlock := container.NewVBox(storageCategories, container.NewHBox(applyButton, refreshButton, a.freezeButton))
	runBlock.Hide()

	startButton.OnTapped = func() {
		storageAddress, err := hexStringToUint64(a.storageAddressEntry.Text)
		if err != nil {
			a.DisplayError(fmt.Errorf("invalid storage address: %v", err))
			return
		}
		t, err := trainer.NewTrainer(storageAddress - a.storageAddressOffset*4)
		if err != nil {
			a.DisplayError(fmt.Errorf("failed to initialize trainer, is the game running?\nError: %v", err))
			return
		}
		a.trainer = t

		a.refreshStorageValues()
		setupBlock.Hide()
		runBlock.Show()
	}

	a.main.SetContent(container.NewVBox(setupBlock, runBlock))
}

func (a *App) initStorageCategory(name string, items []string) fyne.CanvasObject {
	obj := make([]fyne.CanvasObject, len(items))
	categoryCheckbox := widget.NewCheck(name, func(b bool) {
		for _, item := range items {
			index := trainer.StorageIndexes[item]
			a.resourceFreezeChecks[index].SetChecked(b)
		}
	})
	for i, item := range items {
		index := trainer.StorageIndexes[item]
		a.resourceValues[index] = binding.NewInt()
		entry := widget.NewEntryWithData(binding.IntToString(a.resourceValues[index]))
		check := widget.NewCheck(minimumLengthString(item, minResourceLabelLength), nil)
		obj[i] = container.NewVBox(check, entry)
		a.resourceFreezeChecks[index] = check
	}
	return container.NewVBox(categoryCheckbox, newBorder(container.NewVBox(obj...)))
}

func (a *App) refreshStorageValues() {
	if a.trainer == nil {
		return
	}

	values, err := a.trainer.ReadStorageValues()
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

func (a *App) applyStorageValues() {
	if a.trainer == nil {
		return
	}
	values := make([]int32, len(a.resourceValues))
	for i := range a.resourceValues {
		v, err := a.resourceValues[i].Get()
		if err != nil {
			a.DisplayError(fmt.Errorf("failed to get value for resource index %d: %v", i, err))
			return
		}
		values[i] = int32(v)
	}
	err := a.trainer.WriteStorageValues(values)
	if err != nil {
		a.DisplayError(fmt.Errorf("failed to write values to game memory: %v", err))
		return
	}
}

func (a *App) freezeSelectedValues() {
	if a.trainer == nil {
		return
	}

	values := make([]trainer.IndexedValue, 0, len(a.resourceValues))
	for i, check := range a.resourceFreezeChecks {
		if check.Checked {
			v, err := a.resourceValues[i].Get()
			if err != nil {
				a.DisplayError(fmt.Errorf("failed to get value for resource index %d: %v", i, err))
				return
			}
			values = append(values, trainer.IndexedValue{Index: uint64(i), Value: int32(v)})
		}
	}
	if len(values) == 0 {
		a.DisplayError(fmt.Errorf("no values selected to freeze"))
		return
	}
	a.trainer.FreezeStorageValues(values)

	a.freezeButton.Text = freezeButtonTextUnfreeze
	a.freezeButton.OnTapped = a.unfreezeValues
}

func (a *App) unfreezeValues() {
	if a.trainer == nil {
		return
	}
	a.trainer.UnfreezeStorageValues()

	a.freezeButton.Text = freezeButtonTextFreeze
	a.freezeButton.OnTapped = a.freezeSelectedValues
}
