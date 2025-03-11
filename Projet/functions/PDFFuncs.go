package functions

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
)

const tmpDirPath = "tmp"

// TemplateToHTML converts a template to an HTML file, and returns the path to the HTML file and its filename.
func TemplateToHTML(tmp *template.Template, cssPath string, data any) (string, string) {
	filename := GenerateHexFilename()
	path := fmt.Sprintf("%s/%s.html", tmpDirPath, filename)

	// Create a temporary file
	file, err := os.Create(path)
	if err != nil {
		ErrorPrintf("An error occurred while trying to create the temporary file -> %v\n", err)
		return "", ""
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			ErrorPrintf("An error occurred while trying to close the temporary file -> %v\n", err)
		}
	}(file)

	// Execute the template and write the result to the temporary file
	err = tmp.Execute(file, data)
	if err != nil {
		ErrorPrintf("An error occurred while trying to execute the template -> %v\n", err)
		return "", ""
	}
	InfoPrintf("Created HTML file -> %v\n", path)
	InjectCSSIntoHTML(path, cssPath)
	return path, filename
}

// HTMLToPDF converts an HTML file to a PDF file, and returns the path to the PDF file.
func HTMLToPDF(htmlPath string, PDFname string) string {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var pdfBuf []byte

	// Create a temporary file
	pdfPath := fmt.Sprintf("%s/%s.pdf", tmpDirPath, PDFname)
	err := chromedp.Run(ctx,
		chromedp.Navigate(getFileURL(htmlPath)),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			pdfBuf, _, err = page.PrintToPDF().Do(ctx)
			return err
		}),
	)

	if err != nil {
		ErrorPrintf("An error occurred while trying to convert HTML to PDF -> %v\n", err)
		return ""
	}
	// Write the PDF content to a file
	err = os.WriteFile(pdfPath, pdfBuf, 0644)
	if err != nil {
		ErrorPrintf("An error occurred while trying to write the PDF file -> %v\n", err)
		return ""
	}
	InfoPrintf("Created PDF file -> %v\n", pdfPath)
	return pdfPath
}

// TemplateToPDF converts a template to a PDF file, and returns the path to the PDF file.
func TemplateToPDF(tmp *template.Template, cssPath string, data any) string {
	htmlPath, fileName := TemplateToHTML(tmp, cssPath, data)
	if htmlPath == "" || fileName == "" {
		return ""
	}
	return HTMLToPDF(htmlPath, fileName)
}

// InjectCSSIntoHTML modifies the given HTML file by injecting CSS from a CSS file.
func InjectCSSIntoHTML(htmlFilePath, cssFilePath string) {
	// Read the HTML file
	htmlContent, err := os.ReadFile(htmlFilePath)
	if err != nil {
		ErrorPrintf("An error occurred while trying to read the HTML file -> %v\n", err)
		return
	}

	// Read the CSS file
	cssContent, err := os.ReadFile(cssFilePath)
	if err != nil {
		ErrorPrintf("An error occurred while trying to read the CSS file: %v", err)
		return
	}

	// Convert to strings
	htmlStr := string(htmlContent)
	cssStr := string(cssContent)

	// Create the <style> tag with the CSS content
	styleTag := fmt.Sprintf("<style>\n%s\n</style>", cssStr)

	// Regular expressions to locate <head> and existing <style> tags
	headRegex := regexp.MustCompile(`(?i)<head[^>]*>`)
	styleRegex := regexp.MustCompile(`(?i)<style[^>]*>.*?</style>`)

	if styleRegex.MatchString(htmlStr) {
		// Replace existing <style> tag
		htmlStr = styleRegex.ReplaceAllString(htmlStr, styleTag)
	} else if headRegex.MatchString(htmlStr) {
		// Insert CSS right after <head> if no <style> tag exists
		htmlStr = headRegex.ReplaceAllString(htmlStr, "$0\n"+styleTag)
	} else {
		// If no <head> tag exists, insert CSS at the beginning of the file
		htmlStr = styleTag + "\n" + htmlStr
	}

	// Overwrite the original HTML file
	err = os.WriteFile(htmlFilePath, []byte(htmlStr), 0644)
	if err != nil {
		ErrorPrintf("An error occurred while trying to write the HTML file -> %v\n", err)
		return
	}

	InfoPrintf("Injected CSS into HTML file -> %v\n", htmlFilePath)
}

// getFileURL returns the file URL of the given file path.
func getFileURL(filePath string) string {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return ""
	}
	return "file://" + absPath
}

// GenerateHexFilename creates a random hexadecimal filename.
func GenerateHexFilename() string {
	// Generate 6 random bytes and convert them to hexadecimal
	bytes := make([]byte, 6)
	_, err := rand.Read(bytes)
	if err != nil {
		panic("Erreur de génération de nombres aléatoires")
	}

	// Convert to hexadecimal and return
	return hex.EncodeToString(bytes)
}

// MkTempDir creates a temporary directory.
func MkTempDir() {
	// Check if the temporary directory already exists
	if _, err := os.Stat(tmpDirPath); !os.IsNotExist(err) {
		WarningPrintf("Temporary directory already exists -> %v\n", tmpDirPath)
		return
	}

	// Create the temporary directory
	err := os.Mkdir(tmpDirPath, 0755)
	if err != nil {
		ErrorPrintf("An error occurred while trying to create the temporary directory -> %v\n", err)
		return
	}
	InfoPrintf("Created temporary directory -> %v\n", tmpDirPath)
}

// RmTempDir remove all temporary files.
func RmTempDir() {
	// Suppress the temporary directory and all its content
	err := os.RemoveAll(tmpDirPath)
	if err != nil {
		ErrorPrintf("An error occurred while trying to clear the temporary directory -> %v\n", err)
		return
	}
	InfoPrintf("Cleared temporary directory -> %v\n", tmpDirPath)
}
