# Markdown Editor with Live Preview (Gio)

This is a simple Markdown editor built using the Gio GUI toolkit in Go. It allows users to write Markdown in a text editor pane and see a live HTML preview in the adjacent pane.

## Features

- **Markdown Editor**: A simple text editor for writing Markdown.
- **Live Preview**: The content is rendered as HTML in real-time.
- **Cross-Platform**: Works on Windows, Linux, and macOS thanks to Gio.

## Setup Instructions

### Prerequisites

- **Go**: Install Go from [here](https://golang.org/dl/).

### Steps

1. Clone the repository:

```bash
   git clone https://github.com/yourusername/markdown-editor
   cd markdown-editor
```

2. Install the dependencies:

```bash
   go mod tidy
```

3. Run the application:

```bash
   go run .
```

4. Build the application:

```bash
   go build -o markdown-editor .
```
