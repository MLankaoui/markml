# MARKML

<p align="center">
  <a href="https://github.com/MLankaoui/markml"><img src="https://www.canva.com/design/DAGNj_LQ110/TQhnsqXm8j3MuoVOPW3mHg/view?utm_content=DAGNj_LQ110&utm_campaign=designshare&utm_medium=link&utm_source=editor" alt="markml"></a>
</p>

**MARKML** is a simple Markdown to HTML converter designed to transform Markdown syntax into HTML. This tool provides basic support for common Markdown elements including headers, text decorations, links, and images.

## Features

- **Headers**: Converts Markdown headers (`# Header`, `## Header`, etc.) to corresponding HTML header tags (`<h1>`, `<h2>`, etc.).
- **Text Decorations**: Converts Markdown text decorations:
  - `**bold**` to `<bold>bold</bold>`
  - `*italic*` to `<em>italic</em>`
  - `~~strikethrough~~` to `<s>strikethrough</s>`
- **Links**: Converts Markdown links `[Link text](URL)` to HTML anchor tags `<a href='URL'>Link text</a>`.
- **Images**: Converts Markdown images `![alt text](URL)` to HTML image tags `<img src='URL' alt='alt text'>`.

## Installation

To use **MARKML**, you need to have Go installed on your machine. Follow these steps to get started:

1. Clone the repository:
   ```bash
   git clone https://github.com/MLankaoui/markml.git
   cd markml

## Build the project:

```bash
go build -o markml
```
- Run the program:

```bash
./markml
```
# Usage
After building the project, you can run the markml executable to test the converter. The main.go file contains examples of how the Markdown to HTML conversion works. You can modify this file to include your own Markdown text and see the HTML output.

## Testing
To ensure the correctness of the conversion functions, unit tests are provided. Run the tests using:

```bash
go test ./...
```
Contributing
Contributions are welcome! If you have suggestions or improvements, please open an issue or submit a pull request.

## License
This project is licensed under the MIT License. See the LICENSE file for more details.

## Contact
For any questions or feedback, please contact marouanelankaoui66@gmail.com.

Enjoy converting Markdown to HTML with MARKML!