# Website Health Checker

## About

The **Website Health Checker** is a simple command-line tool created in Go to check the health (online or offline status) of a list of websites. This tool also allows for continuous monitoring of websites, updating every 10 seconds, and displays the status in a formatted table. It is designed to automatically handle URLs and adds `https://` if no protocol is provided by the user.

I created this project to learn more about Go programming, working with external libraries, and building interactive CLI applications.

## Features

- **Check multiple websites**: Allows the user to check the status of multiple websites at once.
- **Offline/Online Status**: Displays whether each website is online (`ONLINE ✅`) or offline (`OFFLINE ❌`).
- **Response Time**: Shows how long it took for the website to respond.
- **Continuous Monitoring**: Optionally enables continuous monitoring of the websites, updating the status every 10 seconds.
- **Automatic Protocol Handling**: If a website URL doesn't include the protocol (`https://`), it is automatically added.
- **Formatted Output**: Displays the results in a clean, formatted table.

## What I Learned

- **Go Programming**: How to build and structure a Go project, including working with HTTP requests and response times.
- **Libraries**: How to integrate and work with external libraries such as [huh](https://github.com/charmbracelet/huh) for interactive user input, [lipgloss](https://github.com/charmbracelet/lipgloss) for styled terminal output, and [go-pretty](https://github.com/jedib0t/go-pretty) for formatted table output.
- **CLI Development**: Creating a user-friendly command-line interface that handles input/output and error handling effectively.
- **Concurrency**: Learned how to use Go's concurrency features to handle continuous website monitoring.

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/your-username/website-health-checker.git
   cd website-health-checker

## Installation

1. **Install Go (if not already installed)**:
   Follow the installation instructions from Go's [official website](https://golang.org/doc/install).

2. **Build the project**:
   ```bash
   go build

## Run the application

After building, you can run the program:
./website-health-checker

- **Enter Website URLs**: You will be prompted to enter a list of URLs (comma-separated). You can enter them in the format `https://google.com`, `https://github.com`, or just `google.com` (the protocol will be added automatically).

- **Enable Continuous Monitoring**: You will be asked if you want to enable continuous monitoring. If yes, the program will check the websites every 10 seconds and display the updated status in a table. To stop monitoring, press `Ctrl+C`.

## View Status

The program will display a table with the following columns:

- `#`: The index of the website.
- `Website`: The website URL.
- `Status`: Whether the website is online (ONLINE ✅) or offline (OFFLINE ❌).
- `Response Time`: The time it took for the website to respond.

## Example

When you run the program, it might look something like this:

Enable continuous monitoring? (Yes/No): Yes

🔄 Monitoring enabled... Checking every 10 seconds.

🌐 Website Health Checker
-------------------------------------------
| # | Website            | Status   | Response Time |
-------------------------------------------
| 1 | https://google.com  | ONLINE ✅ | 134ms         |
| 2 | https://github.com  | ONLINE ✅ | 200ms         |
-------------------------------------------


The program will continue checking every 10 seconds and update the table with the latest status and response times.

## Dependencies

 - huh: Used for interactive user input.
 - lipgloss: For styled text output.
 - go-pretty: For displaying the results in a formatted table.

## License

Open Source, feel Free to use as you desire

## Contribution

Feel free to fork this repository, open issues, and submit pull requests. Contributions are welcome!

## Author

Emin Gani