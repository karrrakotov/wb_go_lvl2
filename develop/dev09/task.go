package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"

	"golang.org/x/net/html"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type wgetOptions struct {
	URL       string
	OutputDir string
}

func main() {
	options := parseCommandLineArguments()

	err := download(options.URL, options.OutputDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при скачивании: %v\n", err)
		os.Exit(1)
	}
}

func parseCommandLineArguments() wgetOptions {
	options := wgetOptions{}

	flag.StringVar(&options.URL, "url", "", "URL сайта для скачивания")
	flag.StringVar(&options.OutputDir, "output", ".", "Директория для сохранения файлов")

	flag.Parse()

	return options
}

func download(urlStr string, outputDir string) error {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	err = downloadPage(urlStr, parsedURL, outputDir)
	if err != nil {
		return err
	}

	fmt.Printf("Сайт успешно скачан в директорию: %s\n", outputDir)
	return nil
}

func downloadPage(urlStr string, baseURL *url.URL, outputDir string) error {
	resp, err := http.Get(urlStr)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("не удалось скачать страницу, статус: %d", resp.StatusCode)
	}

	tokenizer := html.NewTokenizer(resp.Body)

	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			return nil
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()

			for _, attr := range token.Attr {
				if attr.Key == "href" || attr.Key == "src" {
					linkURL, err := baseURL.Parse(attr.Val)
					if err != nil {
						fmt.Fprintf(os.Stderr, "Ошибка при анализе URL: %v\n", err)
						continue
					}

					err = downloadResource(linkURL, outputDir)
					if err != nil {
						fmt.Fprintf(os.Stderr, "Ошибка при скачивании ресурса: %v\n", err)
					}
				}
			}
		}
	}
}

func downloadResource(url *url.URL, outputDir string) error {
	resp, err := http.Get(url.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("не удалось скачать ресурс %s, статус: %d", url.String(), resp.StatusCode)
	}

	filePath := path.Join(outputDir, url.Host, url.Path)
	if err := os.MkdirAll(path.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Скачан ресурс: %s\n", filePath)
	return nil
}
