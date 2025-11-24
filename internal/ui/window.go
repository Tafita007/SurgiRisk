package ui

import (
    "github.com/therecipe/qt/widgets"
    "github.com/therecipe/qt/core"
)

type MainWindow struct {
    *widgets.QMainWindow
}

func NewMainWindow() *MainWindow {

    w := &MainWindow{widgets.NewQMainWindow(nil, 0)}
    w.SetWindowTitle("SurgiRisk")

    central := widgets.NewQWidget(nil, 0)
    layout := widgets.NewQVBoxLayout2(central)
    w.SetCentralWidget(central)

    // GroupBox
    card := widgets.NewQGroupBox2("Analyse du Risque Opératoire", nil)
    cardLayout := widgets.NewQVBoxLayout2(card)

    title := widgets.NewQLabel2("Bienvenue dans SurgiRisk", nil, 0)
    title.SetAlignment(core.Qt__AlignCenter)  // ✔ ICI

    input := widgets.NewQLineEdit(nil)
    input.SetPlaceholderText("Entrer l'identifiant du patient...")

    btn := widgets.NewQPushButton2("Analyser", nil)

    cardLayout.AddWidget(title, 0, 0)
    cardLayout.AddWidget(input, 0, 0)
    cardLayout.AddWidget(btn, 0, 0)

    layout.AddWidget(card, 0, 0)

    return w
}
