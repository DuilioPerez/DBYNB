package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func getArticleActivity(title string, image, document *fyne.StaticResource) (aboutLayout *fyne.Container) {
	titleText := widget.NewLabel(title)
	titleText.Wrapping = fyne.TextWrapWord
	titleText.TextStyle.Bold = true
	titleText.Alignment = fyne.TextAlignCenter
	titleText.Refresh()
	mdText := widget.NewRichTextFromMarkdown(string(document.Content()))
	mdText.Wrapping = fyne.TextWrapWord
	mdText.Refresh()
	imageResource := canvas.NewImageFromResource(image)
	imageResource.Resize(fyne.NewSize(200, 200))
	imageResource.FillMode = canvas.ImageFillOriginal
	mainSection := container.NewVBox(titleText,
		widget.NewSeparator(),
		imageResource,
		widget.NewSeparator(), mdText)
	scrollableSection := container.NewVScroll(mainSection)
	scrollableSection.SetMinSize(fyne.NewSize(200, 200))
	iconsSection := container.NewHBox(
		container.NewHBox(),
		widget.NewButtonWithIcon("", theme.HomeIcon(), func() {
			window.SetContent(getMainActivity())
		}),
		widget.NewButtonWithIcon("", theme.InfoIcon(), func() {
			window.SetContent(getAboutActivity())
		}),
		container.NewHBox())
	aboutLayout = container.NewBorder(nil, iconsSection, nil, nil, scrollableSection)
	return
}
