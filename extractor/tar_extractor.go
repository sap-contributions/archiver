package extractor

import (
	"fmt"
	"os"
	"os/exec"
)

type tarExtractor struct{}

func NewTar() Extractor {
	return &tarExtractor{}
}

func (e *tarExtractor) Extract(src, dest string) error {
	if fi, err := os.Stat(dest); err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(dest, 0755); err != nil {
			return fmt.Errorf("failed to create destination directory: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("failed to stat destination: %w", err)
	} else if !fi.IsDir() {
		return fmt.Errorf("destination is not a directory: %s", dest)
	}

	out, err := exec.Command("tar", "xf", src, "-C", dest, "--same-owner").CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to extract tar file: %w\nOutput: %s", err, out)
	}

	return nil
}
