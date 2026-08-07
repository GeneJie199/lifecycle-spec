package redaction

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// Forbidden patterns that must not appear in committed examples or schemas.
// Matches are case-insensitive where appropriate.
var forbidden = []*regexp.Regexp{
	regexp.MustCompile(`(?i)-----BEGIN (RSA |EC |OPENSSH |DSA )?PRIVATE KEY-----`),
	regexp.MustCompile(`(?i)BEGIN RSA PRIVATE KEY`),
	regexp.MustCompile(`(?i)password\s*=`),
	regexp.MustCompile(`(?i)aws_secret`),
	regexp.MustCompile(`ghp_[A-Za-z0-9_]{20,}`),
	regexp.MustCompile(`gho_[A-Za-z0-9_]{20,}`),
}

func repoRoot(t *testing.T) string {
	t.Helper()
	dir, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	return dir
}

func scanText(content string) []string {
	var hits []string
	for _, re := range forbidden {
		if loc := re.FindString(content); loc != "" {
			hits = append(hits, loc)
		}
	}
	return hits
}

func walkJSONFiles(t *testing.T, root string) []string {
	t.Helper()
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.EqualFold(filepath.Ext(path), ".json") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return files
}

func TestExamplesAndSchemasHaveNoSecrets(t *testing.T) {
	root := repoRoot(t)
	dirs := []string{
		filepath.Join(root, "examples"),
		filepath.Join(root, "schemas"),
	}
	var checked int
	for _, dir := range dirs {
		for _, path := range walkJSONFiles(t, dir) {
			raw, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("read %s: %v", path, err)
			}
			hits := scanText(string(raw))
			if len(hits) > 0 {
				t.Errorf("%s matched forbidden patterns: %v", path, hits)
			}
			checked++
		}
	}
	if checked == 0 {
		t.Fatal("no JSON files scanned")
	}
}

func TestScannerMatchesCraftedSecrets(t *testing.T) {
	cases := []string{
		"-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEA...",
		"BEGIN RSA PRIVATE KEY",
		"password=super-secret-value",
		"aws_secret_access_key=FAKE",
		"token ghp_abcdefghijklmnopqrstuvwxyz0123456789",
		"auth gho_abcdefghijklmnopqrstuvwxyz0123456789",
	}
	for _, c := range cases {
		c := c
		t.Run(c[:min(32, len(c))], func(t *testing.T) {
			hits := scanText(c)
			if len(hits) == 0 {
				t.Fatalf("expected scanner to match crafted secret: %q", c)
			}
		})
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
