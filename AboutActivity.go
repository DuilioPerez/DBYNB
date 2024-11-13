package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func getAboutActivity() (aboutLayout *fyne.Container) {
	aboutText := canvas.NewText("Acerca De", color.NRGBA{0, 0, 255, 255})
	aboutText.TextSize = 20
	aboutText.TextStyle.Bold = true
	aboutText.Refresh()
	mdText := widget.NewRichTextFromMarkdown(string(resourceAboutMd.Content()))
	mdText.Wrapping = fyne.TextWrapWord
	mdText.Refresh()
	mainSection := container.NewVBox(aboutText,
		widget.NewSeparator(), mdText)
	scrollableSection := container.NewVScroll(mainSection)
	scrollableSection.SetMinSize(fyne.NewSize(200, 200))
	iconsSection := container.NewHBox(
		container.NewHBox(),
		widget.NewButtonWithIcon("", theme.HomeIcon(), func() {
			window.SetContent(getMainActivity())
		}),
		container.NewHBox())
	aboutLayout = container.NewBorder(nil, iconsSection, nil, nil, scrollableSection)
	return
}
