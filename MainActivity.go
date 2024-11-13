package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func getMainActivity() (mainLayout *fyne.Container) {
	// Articles label.
	articles := canvas.NewText("Artículos", color.NRGBA{0, 0, 255, 255})
	articles.TextSize = 30
	articles.TextStyle.Bold = true
	articles.Refresh()
	// Suggested articles label.
	suggetedArticles := canvas.NewText("Artículos Sugeridos", color.NRGBA{0, 0, 255, 255})
	suggetedArticles.TextSize = 30
	suggetedArticles.TextStyle.Bold = true
	suggetedArticles.Refresh()
	// Phrase label.
	phrase := canvas.NewText("\"El mundo no nos fue dado para llenarlo de desechos; es nuestra responsabilidad cuidar de él, promoviendo lo biodegradable y minimizando lo que no lo es.\"", color.NRGBA{0, 0, 255, 255})
	phrase.TextSize = 10
	phrase.TextStyle.Bold = true
	phrase.Alignment = fyne.TextAlignTrailing
	phrase.Refresh()
	// Cards section.
	cardsSection := container.NewHBox(
		widget.NewCard("Artículos para el lector", "Todos los artículos",
			container.NewVBox(
				widget.NewLabel("Aquí hemos escrito todos los artículos"),
				widget.NewLabel("- El equipo de escritores."))),
		widget.NewCard("Desechos biodegradables", "Duilio Pérez",
			widget.NewButton("Leer", func() {
				window.SetContent(getArticleActivity("Desechos biodegradables",
					resourceDBPng, resourceDBMd))
			})),
		widget.NewCard("Desechos no biodegradables", "Duilio Pérez",
			widget.NewButton("Leer", func() {
				window.SetContent(getArticleActivity("Desechos no biodegradables",
					resourceDNBPng, resourceDNBMd))
			})),
		widget.NewCard("Contaminación por plásticos", "Jesuá Jadhai Rivas Medina",
			widget.NewButton("Leer",
				func() {
					window.SetContent(getArticleActivity("Reciclaje de desechos no biodegradables",
						resourceRDDNBPng, resourceRDDNBMd))
				})),
		widget.NewCard("Gestión de Residuos Sólidos Urbanos", "Jesuá Jadhai Rivas Medina",
			widget.NewButton("Leer",
				func() {
					window.SetContent(getArticleActivity("Gestión de residuos sólidos urbanos",
						resourceGDRSUPng, resourceGDRSUMd))
				})),
		widget.NewCard("Contaminación por Plásticos y Microplásticos",
			"Jesuá Jadhai Rivas Medina",
			widget.NewButton("Leer",
				func() {
					window.SetContent(getArticleActivity(
						"Contaminación por Plásticos y Microplásticos",
						resourceCPPYMPng, resourceCPPYMMd))
				})),
		widget.NewCard("Educación Ambiental para la Reducción de Residuos",
			"Obed Sánchez Domínguez",
			widget.NewButton("Leer",
				func() {
					window.SetContent(getArticleActivity(
						"Educación Ambiental para la Reducción de Residuos",
						resourceEARRPng, resourceEARRMd))
				})),
		widget.NewCard("Producción de Biogás a partir de Desechos Biodegradables",
			"Obed Sánchez Domínguez",
			widget.NewButton("Leer",
				func() {
					window.SetContent(getArticleActivity(
						"Producción de Biogás a partir de Desechos Biodegradables",
						resourcePBADBPng, resourcePBADBMd))
				})),
		widget.NewCard("Estrategia de Reducción de Residuos en Zonas Urbanas y Rurales",
			"Obed Sánchez Domínguez",
			widget.NewButton("Leer",
				func() {
					window.SetContent(getArticleActivity(
						"Estrategia de Reducción de Residuos en Zonas Urbanas y Rurales",
						resourceERRZURPng, resourceERRZURMd))
				})),
		widget.NewCard("Tecnología de Degradación Acelerada para Desechos no Biodegradables",
			"Alex Enrrique Álvarez García",
			widget.NewButton("Leer",
				func() {
					window.SetContent(getArticleActivity(
						"Tecnología de Degradación Acelerada para Desechos No Biodegradables",
						resourceTDADNBPng, resourceTDADNBMd))
				})),
		widget.NewCard("Impacto en la Salud Pública de los Desechos Acumulados",
			"Alex Enrrique Álvarez García",
			widget.NewButton("Leer",
				func() {
					window.SetContent(getArticleActivity(
						"Impacto en la Salud Pública de los Desechos Acumuladosoi",
						resourceISPDAPng, resourceISPDAMd))
				})),
		widget.NewCard("Alternativas Biodegradables a Plásticos Convencionales",
			"Alex Enrrique Álvarez García",
			widget.NewButton("Leer",
				func() {
					window.SetContent(getArticleActivity(
						"Alternativas Biodegradables a Plásticos Convencionales",
						resourceABPCPng, resourceABPCMd))
				})),
	)
	// Suggested section.
	suggestedSection := container.NewHBox(
		widget.NewCard("Artículos sugeridos para el lector", "Mejores artículos",
			container.NewVBox(
				widget.NewLabel("Aquí están los artículos más informativos"),
				widget.NewLabel("- El equipo de escritores."))),
		widget.NewCard("Desechos biodegradables", "Duilio Pérez",
			widget.NewButton("Leer", func() {
				window.SetContent(getArticleActivity("Desechos biodegradables",
					resourceDBPng, resourceDBMd))
			})),
		widget.NewCard("Contaminación por plásticos", "Jesuá Jadhai Rivas Medina",
			widget.NewButton("Leer",
				func() {
					window.SetContent(getArticleActivity("Reciclaje de desechos no biodegradables",
						resourceRDDNBPng, resourceRDDNBMd))
				})),
		widget.NewCard("Producción de Biogás a partir de Residuos Biodegradables", "Obed Sánchez Domínguez",
			widget.NewButton("Leer",
				func() {
					window.SetContent(getArticleActivity("Reciclaje de desechos no biodegradables",
						resourcePBADBPng, resourcePBADBMd))
				})),
		widget.NewCard("Tecnología de Degradación Acelerada para Desechos No Biodegradables",
			"Alex Enrrique Álvarez García",
			widget.NewButton("Leer",
				func() {
					window.SetContent(getArticleActivity(
						"Tecnología de Degradación Acelerada para Desechos No Biodegradables",
						resourceTDADNBPng, resourceTDADNBMd))
				})),
	)
	// Central section.
	centralSection := container.NewVBox(
		articles,
		container.NewHScroll(cardsSection),
		suggetedArticles,
		container.NewHScroll(suggestedSection),
		container.NewHScroll(phrase))
	// Fixed icons container.
	iconsSection := container.NewHBox(
		container.NewHBox(),
		widget.NewButtonWithIcon("", theme.InfoIcon(), func() {
			window.SetContent(getAboutActivity())
		}),
		container.NewHBox())
	// Scrollable section.
	scrollableSection := container.NewVScroll(centralSection)
	// Add the scrollable layout.
	mainLayout = container.NewBorder(nil, iconsSection,
		nil, nil, scrollableSection)
	return
}
