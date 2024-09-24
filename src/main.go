package main

import (
	"log"
	"md-editor-go/src/editor"
	"md-editor-go/src/preview"

	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func main() {
	// Initialize the editor and preview panes
	editor := editor.NewEditorPane()
	prev := preview.NewPreviewPane()

	go func() {
		window := new(app.Window)
		window.Option(app.Title("Markdown Editor with Live Preview"), app.Size(unit.Dp(800), unit.Dp(600)))

		theme := material.NewTheme()
		var ops op.Ops
		for {
			switch e := window.Event().(type) {
			case app.DestroyEvent:
				log.Fatal(e.Err)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				// Layout for the two panes
				layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
					layout.Flexed(2.5, func(gtx layout.Context) layout.Dimensions {
						return editor.Layout(gtx, theme)
					}),
					layout.Flexed(2.5, func(gtx layout.Context) layout.Dimensions {
						return prev.Layout(gtx, theme)
					}),
				)
				e.Frame(gtx.Ops)
			case key.Event:
				// Handle keyboard events here
				editor.HandleKey(e, prev)
			}
		}
	}()
	app.Main()
}
