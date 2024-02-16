package main

import (
    "bufio"
    "fmt"
    "net/http"
    "os"
    "strings"
)

// fetchTLDs retrieves the list of TLDs from the IANA website.
func fetchTLDs(url string) (map[string]bool, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    validTLDs := make(map[string]bool)
    scanner := bufio.NewScanner(resp.Body)
    for scanner.Scan() {
        tld := strings.ToLower(scanner.Text())
        // Skip comments and empty lines
        if strings.HasPrefix(tld, "#") || tld == "" {
            continue
        }
        validTLDs[tld] = true
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return validTLDs, nil
}

func main() {
    tldURL := "https://data.iana.org/TLD/tlds-alpha-by-domain.txt"
    validTLDs, err := fetchTLDs(tldURL)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to fetch TLDs: %v\n", err)
        return
    }

    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        domain := scanner.Text()
        parts := strings.Split(domain, ".")
        if len(parts) < 2 {
            continue // Skip invalid domains
        }

        tld := strings.ToLower(parts[len(parts)-1])
        if _, exists := validTLDs[tld]; exists {
            fmt.Println(domain)
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "reading standard input: %v", err)
    }
}
