package gui

import (
	g "github.com/AllenDang/giu"
	"github.com/Kondou727/maimai-stats-tracker/internal/config"
)

func Loop() {
	g.MainMenuBar().Layout(
		g.Menu("File").Layout(
			g.MenuItem("Open"),
			g.Separator(),
			g.MenuItem("Exit"),
		),
		g.Menu("Misc").Layout(
			g.Button("Button"),
		),
	).Build()
}

func Run(cfg *config.ApiConfig) {

}
