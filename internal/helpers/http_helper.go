package helpers

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadAndExtract(url, dest string) error {
	// Create the file
	tmpFile, err := os.CreateTemp("", "download-*.tar.gz")
	if err != nil {
		return fmt.Errorf("error creating temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up

	// Download the file
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error downloading: %v", err)
	}
	defer res.Body.Close()

	// Write the body to file
	_, err = io.Copy(tmpFile, res.Body)
	if err != nil {
		return fmt.Errorf("error writing to temp file: %v", err)
	}

	// Extract the tar.gz file
	if err := extractTarGz(tmpFile.Name(), dest); err != nil {
		return fmt.Errorf("error extracting: %v", err)
	}

	return nil
}

func extractTarGz(gzipPath, dest string) error {
	gzipFile, err := os.Open(gzipPath)
	if err != nil {
		return err
	}
	defer gzipFile.Close()

	// check if the gzip has already been decompressed and is a tar file
	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)
	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		path := filepath.Join(dest, header.Name)
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(path, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			outFile, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				outFile.Close()
				return err
			}
			outFile.Close()
		}
	}

	return nil
}
