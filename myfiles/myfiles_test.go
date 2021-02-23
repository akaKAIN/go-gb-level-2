package myfiles

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestWriteNewFile_CaseFileExists(t *testing.T) {
	baseData := []byte("Test data")
	tempFile, err := ioutil.TempFile("", "prefix")
	if err != nil {
		t.Fatalf("Creation temp file: %s", err)
	}
	defer func() {
		if err = os.Remove(tempFile.Name()); err != nil {
			t.Fatal("Error of closing temp file for tests", err)
		}
	}()
	err = WriteNewFile(tempFile.Name(), baseData)
	if err == nil {
		t.Fatalf("Must throw error: %s", ErrorFileIsExists)
	}
}

func TestWriteNewFile_CaseCreatingFile(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		data     []byte
	}{
		{
			name:     "Case_1",
			fileName: "fileTestCase1",
			data:     []byte("TestDataCase1"),
		},
		{
			name:     "Case_2",
			fileName: "fileTestCase2",
			data:     []byte(""),
		},
	}

	for _, tc := range tests {
		// Проверяем что не сущетвует одноиенных файлов. Иначе - удаляем их.
		if ok, err := FileIsExists(tc.fileName); ok && err != nil {
			err := RemoveFile(tc.fileName)
			if err != nil {
				t.Fatalf(
					"%s\nCan't remove file: %s, error: %s",
					tc.name,
					tc.fileName,
					ErrorFileIsExists,
				)
			}
		}

		err := WriteNewFile(tc.fileName, tc.data)
		if err != nil {
			t.Fatalf("%s\nmust no errors, but gotten: %s", tc.name, err)
		}
		if ok, err := FileIsExists(tc.fileName); !ok {
			t.Fatalf("%s\ncreated file did not exists. Error: %s", tc.name, err)
		}

		fi, err := os.ReadFile(tc.fileName)
		if err != nil {
			t.Fatalf("%s: %s", tc.name, err)
		}
		if !reflect.DeepEqual(tc.data, fi) {
			t.Fatalf("%s\nMust be %v, but got %v", tc.name, tc.data, fi)
		}

		if ok, err := FileIsExists(tc.fileName); ok {
			if err = os.Remove(tc.fileName); err != nil {
				t.Fatalf(
					"%s\nCan't removed file after success creation: %s",
					tc.name,
					err,
				)
			}
		}
	}
}

func TestRemovingFile(t *testing.T) {
	fileName := fmt.Sprintf("%d", time.Now().UnixNano())
	if _, err := os.Create(fileName); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(fileName); err != nil {
		t.Fatal(err)
	}

	if err := RemoveFile(fileName); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(fileName); err == nil {
		t.Fatalf("File: %s was not removed", fileName)
	}
}
