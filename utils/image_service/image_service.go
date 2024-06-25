package image_service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

type ImageService interface {
    UploadImage(fileHeader *multipart.FileHeader, uploadDir string) (string, error)
    DeleteImage(imageURL string) error
    UpdateImage(existingImageURL string, newFileHeader *multipart.FileHeader, uploadDir string) (string, error)
    GetImageInfo(imageURL string) (os.FileInfo, error)
}

type imageService struct{ }

func NewImageService() ImageService {
    return &imageService{}
}

func (s *imageService) UploadImage(fileHeader *multipart.FileHeader, uploadDir string) (string, error) {
    // Generate a unique file name using UUID
    uniqueId := uuid.New()
    fileName := strings.Replace(uniqueId.String(), "-", "", -1)

    // Ensure the file has an extension and handle it properly
    fileExt := filepath.Ext(fileHeader.Filename)
    if fileExt == "" {
        return "", fmt.Errorf("file has no extension")
    }

    // Generate the full file name with the extension
    image := fmt.Sprintf("%s%s", fileName, fileExt)

    // Ensure the upload directory exists
    if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
        return "", fmt.Errorf("failed to create upload directory: %v", err)
    }

    // Create the destination file path
    destPath := filepath.Join(uploadDir, image)

    // Open the source file
    src, err := fileHeader.Open()
    if err != nil {
        return "", fmt.Errorf("failed to open file: %v", err)
    }
    defer src.Close()

    // Create the destination file
    dst, err := os.Create(destPath)
    if err != nil {
        return "", fmt.Errorf("failed to create destination file: %v", err)
    }
    defer dst.Close()

    // Copy the file content to the destination file
    if _, err := io.Copy(dst, src); err != nil {
        return "", fmt.Errorf("failed to copy file: %v", err)
    }

    // Generate the file URL
    fileURL := fmt.Sprintf("/%s/%s", uploadDir, image)
    return fileURL, nil
}

func (s *imageService) DeleteImage(imageURL string) error {
	fileName := filepath.Base(imageURL)
	filePath := filepath.Join("uploads", fileName)

	if _, err := os.Stat(filePath); err == nil {
		return os.Remove(filePath)
	} else if os.IsNotExist(err) {
		return nil
	} else {
		return err
	}
}

func (s *imageService) UpdateImage(existingImageURL string, newFileHeader *multipart.FileHeader, uploadDir string) (string, error) {
	// Delete the existing image
	if err := s.DeleteImage(existingImageURL); err != nil {
		return "", err
	}

	// Upload the new image
	newImageURL, err := s.UploadImage(newFileHeader, uploadDir)
	if err != nil {
		return "", err
	}

	return newImageURL, nil
}

func (s *imageService) GetImageInfo(imageURL string) (os.FileInfo, error) {
	fileName := filepath.Base(imageURL)
	filePath := filepath.Join("uploads", fileName)

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	return fileInfo, nil
}
