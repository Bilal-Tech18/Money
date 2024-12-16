# Spid3r - Website Crawler

**Spid3r** is a powerful tool designed to crawl websites, extracting every page, input field (`input`, `textarea`), and button (`button`, `a`, `div`, `span` with interactive roles). It offers flexible output options, allowing results to be saved in Excel, text or JSON formats, making it ideal for a variety of applications such as data collection and website analysis.

![spider](https://github.com/user-attachments/assets/db2079af-b30c-4bd2-b439-63a5f1509256)

## Installation

Before using Spid3r, please ensure that all required libraries are installed by running the installation script:

```
./Spid3r_installer.sh
```
## Tool Compilation

Once all dependencies have been installed, compile the tool using the following command:

```
go build crawling.go createExcelFile.go createTxtFile.go saveJson.go main.go
```
## Usage

You can run Spid3r directly with the following command:

```
./crawling [--excel|--txt|--json|--verbose|--help] <url>
```

### Options :

--excel: Saves the output as an Excel file.
--txt: Saves the output as a text file.
--json: Saves the output as a JSON file.
--verbose: Displays each URL visited in the terminal, as well as information on elements found (inputs, buttons, etc.).
--help: Displays usage information.

### Features

Input Element Detection: Identifies all input and textarea elements.

Button Detection:
   - Detects all HTML `<button>` elements.
   - Identifies `<a>` links with `role="button"`.
   - Detects `<div>` and `<span>` elements with `role="button"` or accessibility attributes like `aria-pressed` or `aria-haspopup`.

Output Options: Saves the results in Excel, text, or JSON format for detailed analysis.

### Example Commande:


```
./crawling --json “https://example.com”
```

This command crawls the website and saves the results in a JSON file.
