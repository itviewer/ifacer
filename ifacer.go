package main

import (
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "ifacer/utils"
)

var (
    ifaceOptions []string
)

func init() {
}

func main() {
    a := app.New()
    w := a.NewWindow("ifacer")

    lanRoute := widget.NewMultiLineEntry()
    lanRoute.Append("10.0.0.0/255.0.0.0\n")
    lanRoute.Append("20.0.0.0/255.0.0.0")

    lanInterfaceCombo := widget.NewSelect(ifaceOptions, func(value string) {})
    lanInterfaceCombo.SetSelected("Option 1")

    lanContainer := container.NewVBox(lanRoute, lanInterfaceCombo)
    lanCard := widget.NewCard("Intranet", "", lanContainer)

    wanRoute := widget.NewEntry()
    wanRoute.SetText("0.0.0.0/0.0.0.0")
    wanRoute.Disable()
    wanInterfaceCombo := widget.NewSelect(ifaceOptions, func(value string) {})
    wanInterfaceCombo.SetSelected("Option 2")

    wanContainer := container.NewVBox(wanRoute, wanInterfaceCombo)
    wanCard := widget.NewCard("Internet", "", wanContainer)

    w.SetContent(container.NewVBox(
        lanCard,
        wanCard,
        widget.NewButton("Ok", func() {
            utils.SetRoutes(lanRoute.Text, wanRoute.Text, lanInterfaceCombo.Selected, wanInterfaceCombo.Selected)
        }),
    ))

    w.CenterOnScreen()
    w.ShowAndRun()
}
