# MARKML

<p align="center">
  <a href="https://github.com/MLankaoui/markml"><img src="MARKML.png" alt="markml"></a>
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
## TODO

- **Lists**:
  - **Unordered Lists**: Support Markdown unordered lists (`* item`, `- item`, `+ item`) and convert them to HTML `<ul>` and `<li>` tags.
  - **Ordered Lists**: Support Markdown ordered lists (`1. item`, `2. item`) and convert them to HTML `<ol>` and `<li>` tags.

- **Code Blocks**:
  - **Inline Code**: Convert inline code (`` `code` ``) to HTML `<code>code</code>`.
  - **Fenced Code Blocks**: Convert fenced code blocks (```` ```language \n code \n ``` ````) to HTML `<pre><code>code</code></pre>`.

- **Blockquotes**:
  - Convert Markdown blockquotes (`> quote`) to HTML `<blockquote>quote</blockquote>`.

- **Tables**:
  - Convert Markdown tables (with pipes and dashes) to HTML `<table>`, `<tr>`, `<th>`, and `<td>` tags.

- **Links and Images with Titles**:
  - Support Markdown links with titles (`[text](url "title")`) and images with titles (`![alt text](url "title")`), converting them to HTML `<a href='url' title='title'>text</a>` and `<img src='url' alt='alt text' title='title'>`.

- **HTML Escaping**:
  - Implement escaping for HTML special characters to prevent rendering issues and security risks.

- **Error Handling and Validation**:
  - Add error handling for malformed Markdown inputs to ensure graceful degradation.

- **Custom Extensions and Configuration**:
  - Implement support for custom Markdown extensions and user configuration options.

- **Command-Line Interface Enhancements**:
  - Improve CLI support to accept input files and output options for better user experience.


## Contributing
Contributions are welcome! If you have suggestions or improvements, please open an issue or submit a pull request.

## License
This project is licensed under the MIT License. See the LICENSE file for more details.

## Contact
For any questions or feedback, please contact marouanelankaoui66@gmail.com.

Enjoy converting Markdown to HTML with MARKML!