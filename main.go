package main

import (
	"math/rand"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := app.New()
	w := a.NewWindow("Saigo")
	w.Resize(fyne.NewSize(900, 700))

	messages := widget.NewLabel("Message Bar: No Messages.")

	// The number of characters in each column, separated by spaces
	charCountLabel := widget.NewLabel("Char Count")
	charCount := widget.NewEntry()
	charCount.SetPlaceHolder(charCountLabel.Text)
	charCount.SetText("5")

	// The groups in each row of a message, separated by newlines
	columnCountLabel := widget.NewLabel("Column Count")
	columnCount := widget.NewEntry()
	columnCount.SetPlaceHolder(columnCountLabel.Text)
	columnCount.SetText("15")

	// The number of rows per message
	rowCountLabel := widget.NewLabel("Row Count")
	rowCount := widget.NewEntry()
	rowCount.SetPlaceHolder(rowCountLabel.Text)
	rowCount.SetText("15")

	pad := widget.NewMultiLineEntry()
	pad.SetText("OTP Output Here")
	pad.TextStyle = fyne.TextStyle{Monospace: true}
	pad.Resize(fyne.NewSize(900, 300))

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			messages.SetText("Create Icon: TBD")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {
			messages.SetText("Cut Icon: TBD")
		}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
			messages.SetText("Copy Icon: TBD")
		}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {
			messages.SetText("Paste Icon: TBD")
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			dialog.NewInformation(
				"About Saigo",
				"Saigo Version 0.0.1\n(c) 2022 Robert McAtee",
				w,
			).Show()
		}),
	)

	buttonOtp := widget.NewButton("Generate One Time Pad", func() {
		charCountInt, err := strconv.Atoi(charCount.Text)
		if err != nil {
			// Add code here to handle the error!
		}

		columnCountInt, err := strconv.Atoi(columnCount.Text)
		if err != nil {
			// Add code here to handle the error!
		}

		rowCountInt, err := strconv.Atoi(rowCount.Text)
		if err != nil {
			// Add code here to handle the error!
		}

		padString := ""
		pad.SetText("")
		padString = generatePadString(padString, rowCountInt, columnCountInt, charCountInt)
		pad.SetText(padString)
	})

	w.SetContent(container.NewVBox(
		toolbar,
		messages,
		charCountLabel,
		charCount,
		columnCountLabel,
		columnCount,
		rowCountLabel,
		rowCount,
		buttonOtp,
		container.NewWithoutLayout(pad)))

	w.ShowAndRun()
}

func generatePadString(padString string, rowCountInt int, columnCountInt int, charCountInt int) string {
	for i := 0; i < rowCountInt; i++ {
		for ii := 0; ii < columnCountInt; ii++ {
			for iii := 0; iii < charCountInt; iii++ {
				padString = padString + generateRandomChar()
			}
			padString = padString + " "
		}
		padString = padString + "\n"
	}
	return padString
}

func generateRandomChar() string {
	randomChar := 'A'
	randomInt := rand.Intn(27)
	if randomInt == 1 {
		randomChar = '0' + rune(rand.Intn(10))
	} else {
		randomChar = 'A' + rune(rand.Intn(26))
	}
	return string(randomChar)
}
