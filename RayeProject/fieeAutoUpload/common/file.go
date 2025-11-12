package common

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// zipFolder 将整个文件夹打包成 zipFilePath
func ZipFolder(folderPath, zipFilePath string) error {
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	err = filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(filepath.Dir(folderPath), path) // 相对路径
		if err != nil {
			return err
		}

		fw, err := archive.Create(relPath)
		if err != nil {
			return err
		}

		fr, err := os.Open(path)
		if err != nil {
			return err
		}

		_, err = io.Copy(fw, fr)
		fr.Close()
		return err
	})
	return err
}
func CopyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, _ := filepath.Rel(src, path)
		targetPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(targetPath, info.Mode())
		} else {
			srcFile, err := os.Open(path)
			if err != nil {
				return err
			}
			defer srcFile.Close()

			dstFile, err := os.OpenFile(targetPath, os.O_CREATE|os.O_WRONLY, info.Mode())
			if err != nil {
				return err
			}
			defer dstFile.Close()

			_, err = io.Copy(dstFile, srcFile)
			return err
		}
	})
}
