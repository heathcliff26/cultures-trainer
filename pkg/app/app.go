package app

import (
	"fmt"

	"fyne.io/fyne/v2"
	fApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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
	resourceEntries      []*Int32Entry
	freezeButton         *widget.Button
}

func New() *App {
	app := newApp()
	version := getVersion(app)
	main := app.NewWindow(version.Name)

	a := &App{
		app:             app,
		main:            main,
		version:         version,
		resourceEntries: make([]*Int32Entry, trainer.ResourceCount),
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
		a.initStorageCategory("Nahrung", trainer.CategoryNahrung),
		a.initStorageCategory("Bauwaren", trainer.CategoryBauwaren),
		a.initStorageCategory("Resourcen", trainer.CategoryResourcen),
		a.initStorageCategory("Waffen", trainer.CategoryWaffen),
		a.initStorageCategory("Bonusgegenstände", trainer.CategoryBonus),
		a.initStorageCategory("Sonstiges", trainer.CategorySonstiges),
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
			a.resourceEntries[index].Checkbox.SetChecked(b)
		}
	})
	for i, item := range items {
		index := trainer.StorageIndexes[item]
		a.resourceEntries[index] = NewInt32Entry(item)
		obj[i] = a.resourceEntries[index]
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
	for i := range a.resourceEntries {
		a.resourceEntries[i].Set(values[i])
	}
}

func (a *App) applyStorageValues() {
	if a.trainer == nil {
		return
	}
	values := make([]int32, trainer.ResourceCount)
	for i := range a.resourceEntries {
		values[i] = a.resourceEntries[i].Get()
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

	values := make([]trainer.IndexedValue, 0, trainer.ResourceCount)
	for i, entry := range a.resourceEntries {
		if entry.Checkbox.Checked {
			values = append(values, trainer.IndexedValue{Index: i, Value: a.resourceEntries[i].Get()})
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
