package desktop

import (
	"encoding/json"
	"os"

	"github.com/Kwynto/preserves"
)

func (a *agonistApp) homeBtn() func() {
	return func() {
		a.winElem.settingsCard.Show()
		a.winElem.alphabetCard.Hide()
		a.winElem.outdateCard.Hide()
		a.winElem.genSiteCard.Hide()
		a.winElem.aboutCard.Hide()
	}
}

func (a *agonistApp) alphabetBtn() func() {
	return func() {
		a.winElem.settingsCard.Hide()
		a.winElem.alphabetCard.Show()
		a.winElem.outdateCard.Hide()
		a.winElem.genSiteCard.Hide()
		a.winElem.aboutCard.Hide()
	}
}

func (a *agonistApp) outdateBtn() func() {
	return func() {
		a.winElem.settingsCard.Hide()
		a.winElem.alphabetCard.Hide()
		a.winElem.outdateCard.Show()
		a.winElem.genSiteCard.Hide()
		a.winElem.aboutCard.Hide()
	}
}

func (a *agonistApp) genSiteBtn() func() {
	return func() {
		a.winElem.settingsCard.Hide()
		a.winElem.alphabetCard.Hide()
		a.winElem.outdateCard.Hide()
		a.winElem.genSiteCard.Show()
		a.winElem.aboutCard.Hide()
	}
}

func (a *agonistApp) aboutBtn() func() {
	return func() {
		a.winElem.settingsCard.Hide()
		a.winElem.alphabetCard.Hide()
		a.winElem.outdateCard.Hide()
		a.winElem.genSiteCard.Hide()
		a.winElem.aboutCard.Show()
	}
}

func (a *agonistApp) saveSettings() func() {
	return func() {
		a.env.GhTiket = a.winElem.settToken.Text
		a.env.SourcePath = a.winElem.settSource.Text

		// Loading a target file
		a.outSettLog("README.md downloading started.")
		source := preserves.ConcatBuffer(a.env.SourcePath, "archive/refs/heads/main.zip")
		_, err := preserves.DownloadFile(source, "./data/")
		if err != nil {
			a.outSettLog("Loading README.md failed.")
			return
		}
		a.outSettLog("Finished loading README.md.")

		// TODO: Unzip

		// Save enveroment to JSON
		out, err := os.Create("./data/cfg/settings.json")
		if err != nil {
			return
		}
		encodeJSON := json.NewEncoder(out)
		err = encodeJSON.Encode(a.env)
		if err != nil {
			return
		}
		a.outSettLog("Enveroment saved.")

		a.env.IsReady = true
		a.winElem.alphabetBtn.Show()
		a.winElem.outdateBtn.Show()
		a.winElem.genSiteBtn.Show()
	}
}

func (a *agonistApp) testAlpha() func() {
	return func() {
		//
	}
}

func (a *agonistApp) testOutdate() func() {
	return func() {
		//
	}
}

func (a *agonistApp) generateSite() func() {
	return func() {
		//
	}
}
