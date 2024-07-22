package storage

import (
	"encoding/csv"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"time"

	"github.com/akosgarai/projectregister/pkg/config"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// CSVStorage interface for storing csv files
type CSVStorage interface {
	Save(file multipart.File) (string, error)
	Delete(fileName string) error
	Read(fileName string) ([][]string, error)
}

type storage struct {
	basePath string
}

// generates a random filename
func (s *storage) generateFileName(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(b)
}

// save saves the data to a file
func (s *storage) saveMultipartFile(file multipart.File, extenstion string) (string, error) {
	// generate a unique file name
	fileName := s.generateFileName(64)
	defer file.Close()
	targetFileName := s.basePath + "/" + fileName + "." + extenstion

	// create a new file
	f, err := os.OpenFile(targetFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// copy the file to the target file
	_, err = io.Copy(f, file)
	if err != nil {
		return "", err
	}

	return fileName, nil
}

// deleteFile deletes the file
func (s *storage) deleteFile(fileName string) error {
	targetFileName := s.basePath + "/" + fileName
	return os.Remove(targetFileName)
}

// readFile reads the file
func (s *storage) readFile(fileName string) (string, error) {
	targetFileName := s.basePath + "/" + fileName
	fileContent, err := os.ReadFile(targetFileName)
	if err != nil {
		return "", err
	}
	return string(fileContent), nil
}

// CSVFileStorage type
type CSVFileStorage struct {
	storage
}

// NewCSVFileStorage creates a new CSVFileStorage
func NewCSVFileStorage(envConfig *config.Environment) *CSVFileStorage {
	return &CSVFileStorage{
		storage: storage{
			basePath: envConfig.GetUploadDirectoryPath(),
		},
	}
}

// Save saves the data to a file. Returns the filename and an error
func (s *CSVFileStorage) Save(file multipart.File) (string, error) {
	return s.saveMultipartFile(file, "csv")
}

// Delete deletes the file
func (s *CSVFileStorage) Delete(fileName string) error {
	return s.deleteFile(fileName)
}

// Read reads the file
func (s *CSVFileStorage) Read(fileName string) ([][]string, error) {
	targetFileName := s.basePath + "/" + fileName
	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.File
	file, err := os.Open(targetFileName)
	if err != nil {
		return nil, err
	}

	// close the file
	defer file.Close()

	// The csv.NewReader() function is called in
	// which the object os.File passed as its parameter
	// and this creates a new csv.Reader that reads
	// from the file
	reader := csv.NewReader(file)

	// ReadAll reads all the records from the CSV file
	// and Returns them as slice of slices of string
	// and an error if any
	return reader.ReadAll()
}
