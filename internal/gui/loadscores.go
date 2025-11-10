package gui

import (
	"github.com/AllenDang/giu"
	"github.com/Kondou727/maimai-stats-tracker/internal/config"
)

var tsv string

func LoadScoreView(cfg *config.ApiConfig) giu.Layout {
	return giu.Layout{
		giu.Label("Enter raw TSV:"),
		giu.InputText(&tsv).Label("TSV"),
		giu.Labelf("You typed %s", tsv),
	}
}
