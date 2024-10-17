Spid3r - Website Crawler

Spid3r is a powerful tool designed to crawl websites, extracting every page, input field, and button. It offers flexible output options, allowing you to save results in Excel or text formats, making it ideal for a wide range of applications such as data collection and website analysis.
Installation

Before using Spid3r, ensure that all required libraries are installed by running the installation script:

bash

./Spid3r_installer.sh

Building the Tool

Once the dependencies are installed, compile the tool using the following command:

bash

go build crawling.go createExcelFile.go createTxtFile.go main.go

Usage

You can run Spid3r directly by executing:

bash

go run main.go [--excel|--txt|--verbose|--help] <url>

Options:

    --excel: Save the output to an Excel file.
    --txt: Save the output to a text file.
    --verbose: Print each visited URL in the terminal.
    --help: Display usage information.

Example

bash

go run main.go --excel https://example.com

This will crawl the website and save the results to an Excel file.