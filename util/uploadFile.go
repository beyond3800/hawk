package util

import "os"
import "io"
import "path/filepath"

func UploadFile(file string, destDir string) error {
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		err = os.MkdirAll(destDir, 0755)
		if err != nil {
			return err
		}
	}

	// 2. Open the source file
	srcFile, err := os.Open(file)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// 3. Build destination path
	filename := filepath.Base(file) // extract "image.png" from "/path/image.png"
	destPath := filepath.Join(destDir, filename)

	// 4. Create destination file
	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// 5. Copy content
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}