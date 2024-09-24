package preview

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/russross/blackfriday/v2"
)

type PreviewPane struct {
	content string
}

// NewPreviewPane initializes a new preview pane
func NewPreviewPane() *PreviewPane {
	return &PreviewPane{}
}

// SetContent allows updating the content to be previewed
func (p *PreviewPane) SetContent(markdownText string) {
	p.content = markdownText
}

// Layout renders the Markdown preview
func (p *PreviewPane) Layout(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	// Convert Markdown to HTML (for rendering, we still display it as plain text)
	html := blackfriday.Run([]byte(p.content))
	renderedContent := string(html)

	// Create a label for the rendered Markdown
	label := material.Label(theme, unit.Sp(50), renderedContent)

	// Layout the label
	return label.Layout(gtx)
}
