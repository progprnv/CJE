package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// ANSI color codes for CLI output
const (
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorReset  = "\033[0m"
)

// Function to read subdomains from a file
func readSubdomains(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var subdomains []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		subdomains = append(subdomains, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return subdomains, nil
}

// Function to test Cookie Jar Overflow
func testCookieJarOverflow(url string) string {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}

	// Add excessive cookies to test overflow
	for i := 0; i < 60; i++ {
		req.AddCookie(&http.Cookie{
			Name:  fmt.Sprintf("testCookie%d", i),
			Value: fmt.Sprintf("testValue%d", i),
		})
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	defer resp.Body.Close()

	if len(resp.Cookies()) > 50 {
		return ColorRed + "Possibly Vulnerable: Server accepted too many cookies" + ColorReset
	}
	return ColorGreen + "Not Vulnerable: Proper cookie handling" + ColorReset
}

// Function to test Session Fixation
func testSessionFixation(url string) string {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Set a fixed session cookie
	fixedSessionCookie := &http.Cookie{
		Name:  "sessionID",
		Value: "fixedSessionID",
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	req.AddCookie(fixedSessionCookie)

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	defer resp.Body.Close()

	for _, cookie := range resp.Cookies() {
		if cookie.Name == "sessionID" && cookie.Value == "fixedSessionID" {
			return ColorRed + "Vulnerable: Session fixation possible" + ColorReset
		}
	}
	return ColorGreen + "Not Vulnerable: Session IDs are regenerated" + ColorReset
}

// Main function
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: CookieJarExploiter <subdomain_file>")
		os.Exit(1)
	}

	fileName := os.Args[1]
	subdomains, err := readSubdomains(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Starting Cookie Jar Exploiter...")
	for _, subdomain := range subdomains {
		url := fmt.Sprintf("https://%s", strings.TrimSpace(subdomain))
		fmt.Printf("\nTesting %s...\n", url)

		// Test for Cookie Jar Overflow
		cookieResult := testCookieJarOverflow(url)
		fmt.Printf("  Cookie Jar Overflow: %s\n", cookieResult)

		// Test for Session Fixation
		sessionResult := testSessionFixation(url)
		fmt.Printf("  Session Fixation: %s\n", sessionResult)
	}
}
