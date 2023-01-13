package main

import (
    "github.com/gotk3/gotk3/gtk"
    "log"
)

func main() {
    gtk.Init(nil)

    win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)

    if err != nil {
        log.Panic(err)
    }

    win.Connect("destroy", func() {
        gtk.MainQuit()
    })

    verticalBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)

    if err != nil {
        log.Panic(err)
    }

    win.Add(verticalBox)

    box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 50)

    if err != nil {
        log.Panic(err)
    }

    verticalBox.Add(box)

    button, err := gtk.ButtonNewWithLabel("foo")

    if err != nil {
        log.Panic(err)
    }

    button.SetSizeRequest(100, 100)
    button.SetMarginStart(10)
    button.SetMarginTop(10)

    box.Add(button)

    scrolledWindow, err := gtk.ScrolledWindowNew(nil, nil)

    if err != nil {
        log.Panic(err)
    }

    scrolledWindow.SetMinContentHeight(250)
    scrolledWindow.SetMinContentWidth(250)
    scrolledWindow.SetMaxContentHeight(250)
    scrolledWindow.SetMaxContentWidth(250)

    box.Add(scrolledWindow)

    input, err := gtk.TextViewNew()

    if err != nil {
        log.Panic(err)
    }

    input.SetSizeRequest(100, 100)
    input.SetMarginTop(10)
    input.SetMarginEnd(10)

    scrolledWindow.Add(input)

    errorLabel, err := gtk.LabelNew("")

    if err != nil {
        log.Panic(err)
    }

    verticalBox.Add(errorLabel)

    buf, err := input.GetBuffer()

    if err != nil {
        log.Panic(err)
    }

    buf.SetText(`
button {

}
`)

    provider, err := gtk.CssProviderNew()

    if err != nil {
        log.Panic(err)
    }

    button.Connect("clicked", func(btn *gtk.Button) {
        startIter := buf.GetStartIter()
        endIter := buf.GetEndIter()

        text, err := buf.GetText(startIter, endIter, false)

        if err != nil {
            log.Panic(err)
        }

        styleContext, err := btn.GetStyleContext()

        if err != nil {
            log.Panic(err)
        }

        err = provider.LoadFromData(text)

        if err != nil {
            errorLabel.SetText(err.Error())
            return
        } else {
            errorLabel.SetText("")
        }

        styleContext.RemoveProvider(provider)
        styleContext.AddProvider(provider, gtk.STYLE_PROVIDER_PRIORITY_USER)
    })

    win.ShowAll()

    gtk.Main()
}
