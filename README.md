# Markdown to HTML Converter

This is a simple Go program that converts Markdown files (.md) to HTML file (.html) .

## How It Works

- The program reads a Markdown file.
- It converts the content to HTML.
- It saves the converted HTML in a new file.

## Requirements

- **Go** (version 1.18 or later) needs to be installed.

## Setup

1. **Install Go**: If you don't have Go installed, you can download it from [here](https://golang.org/dl/).

2. **Clone the project**:

   ```bash 
   git clone https://github.com/n1tik91/Markdown---HTML.git
   ```
   ```bash 
   cd Markdown---HTML
   ```

## Install the necessary package:

Run this command to install the `blackfriday` package:

```bash
go get github.com/russross/blackfriday/v2

```


## How to Use

Run the program by typing:

1. ``` bash
   go run main.go 
   ```

2. ``` bash
   go run main.go Markdownfile.md
   ```


The program will create a new `.html` file with the same name as your Markdown file (e.g., `Markdownfile.md.html`).
