package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"syscall"
	"time"

	"os/signal"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/jedib0t/go-pretty/v6/table"
)

// Define styles for online and offline indicators
var green = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00")).Bold(true)
var red = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Bold(true)
var blue = lipgloss.NewStyle().Foreground(lipgloss.Color("#0000FF")).Bold(true)

// Function to check if a website is online
func checkWebsite(url string) (string, time.Duration) {
	start := time.Now() // Start timing the request

	// Ensure the URL includes the protocol (http:// or https://)
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url // Default to https:// if no protocol is specified
	}

	// Create an HTTP client with a user agent
	client := &http.Client{
		Timeout: 10 * time.Second, // Set a timeout for the request
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err) // Debugging error
		return red.Render("OFFLINE ❌"), time.Since(start)
	}

	// Set a common User-Agent header to mimic a real browser request
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")

	resp, err := client.Do(req)
	elapsed := time.Since(start) // Calculate response time

	if err != nil {
		fmt.Println("Request error:", err) // Debugging error
		return red.Render("OFFLINE ❌"), elapsed
	}

	// Log the status code for debugging
	fmt.Println("Status Code for", url, ":", resp.StatusCode)

	if resp.StatusCode >= 400 {
		return red.Render("OFFLINE ❌"), elapsed
	}
	return green.Render("ONLINE ✅"), elapsed
}

func main() {
	var urls string
	var monitor bool

	// User input: Get a list of websites (comma-separated)
	huh.NewInput().
		Title("Enter Website URLs (comma-separated)").
		Placeholder("e.g., google.com, github.com").
		Value(&urls).
		Run()

	// User input: Enable continuous monitoring?
	huh.NewConfirm().
		Title("Enable continuous monitoring?").
		Affirmative("Yes").
		Negative("No").
		Value(&monitor).
		Run()

	// Convert comma-separated string to a slice of URLs
	urlList := strings.Split(strings.TrimSpace(urls), ",")

	// Function to display website status in a table
	displayStatus := func() {
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetTitle("🌐 Website Health Checker")
		t.AppendHeader(table.Row{"#", "Website", "Status", "Response Time"})

		// Loop through all entered URLs
		for i, url := range urlList {
			url = strings.TrimSpace(url) // Remove extra spaces
			status, responseTime := checkWebsite(url)
			t.AppendRow([]interface{}{i + 1, blue.Render(url), status, responseTime})
		}

		t.Render() // Print table
	}

	// Initial status check
	displayStatus()

	// If monitoring is enabled, keep checking every 10 seconds
	if monitor {
		// Create a channel to listen for the interrupt signal
		stopChan := make(chan os.Signal, 1)
		signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

		// Goroutine to monitor continuously
		go func() {
			fmt.Println("\n🔄 Monitoring enabled... Checking every 10 seconds.")
			for {
				select {
				case <-stopChan:
					fmt.Println("\n🚨 Monitoring stopped.")
					return
				default:
					time.Sleep(10 * time.Second)
					fmt.Print("\033[H\033[2J") // Clear screen
					displayStatus()
				}
			}
		}()

		// Wait for interrupt signal to stop monitoring
		<-stopChan
	}
}
