package editor

import (
	"md-editor-go/src/preview"

	"gioui.org/font"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type EditorPane struct {
	editor widget.Editor
}

// NewEditorPane creates a new instance of the Markdown editor
func NewEditorPane() *EditorPane {
	e := widget.Editor{
		SingleLine: false, // Allow multi-line input
		Submit:     true,
	}
	return &EditorPane{
		editor: e,
	}
}

// Layout renders the Markdown editor
func (e *EditorPane) Layout(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	ed := material.Editor(theme, &e.editor, "Write your markdown here...")
	ed.Font = font.Font{Typeface: "Mono"}
	return ed.Layout(gtx)
}

// HandleKey processes key inputs for the editor
func (e *EditorPane) HandleKey(ev key.Event, preview *preview.PreviewPane) {
	if ev.State == key.Press {
        // Check if the key pressed is Enter
        if ev.Name == key.NameReturn {
            e.editor.Insert("\n") // Insert a newline directly
        }
        // Update the preview with the current text
        preview.SetContent(e.GetText())
    }
}

// GetText returns the current content of the editor
func (e *EditorPane) GetText() string {
	return e.editor.Text()
}
