package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

var appAgonist agonistApp

func (a *agonistApp) createElements() {
	a.createCardSettings()
	a.createCardAlphabet()
	a.createCardOutdate()
	a.createCardAbout()

	a.createMenuButtons()
}

func (a *agonistApp) createApplication() {
	a.app = app.New()

	a.mainWindow = a.app.NewWindow("[AGo]nist")
	{
		icon, _ := fyne.LoadResourceFromPath("icon.png")
		a.mainWindow.SetIcon(icon)
	}
	a.mainWindow.Resize(fyne.NewSize(700, 415))
	a.mainWindow.CenterOnScreen()
	a.mainWindow.SetFixedSize(true)

	a.createElements()

	a.workSpace = container.NewGridWrap(fyne.NewSize(650, 405), a.winElem.settingsCard, a.winElem.alphabetCard, a.winElem.outdateCard, a.winElem.aboutCard)
	a.mainMenuBox = container.NewVBox(a.winElem.homeBtn, a.winElem.alphabetBtn, a.winElem.outdateBtn, a.winElem.aboutBtn)
	a.mainHBox = container.NewHBox(a.mainMenuBox, a.workSpace)

	a.mainWindow.SetContent(a.mainHBox)
}

func Run() {
	appAgonist.createApplication()

	appAgonist.winElem.settingsCard.Show()
	appAgonist.winElem.alphabetCard.Hide()
	appAgonist.winElem.outdateCard.Hide()
	appAgonist.winElem.aboutCard.Hide()

	appAgonist.mainWindow.ShowAndRun()
}
