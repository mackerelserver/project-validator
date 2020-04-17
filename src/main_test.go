package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func collectYamlFiles(folder string) []string {
	var files []string
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".yml") || strings.HasSuffix(path, ".yaml") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

func TestValidConfigurations(t *testing.T) {
	Convey("test valid configurations", t, func() {
		for _, file := range collectYamlFiles("../examples/valid") {
			Convey(fmt.Sprintf("file %s should validate", file), func() {
				result, err := ValidateFile(file)
				So(err, ShouldBeNil)
				So(result, ShouldValidate)
			})
		}
	})
}

func ShouldNotValidate(actual interface{}, _ ...interface{}) string {
	result := ShouldValidate(actual)
	if result == "" { // file validated
		return "File validated (it should not)"
	} else {
		return ""
	}
}

func TestInvalidConfigurations(t *testing.T) {
	Convey("test invalid configurations", t, func() {
		for _, file := range collectYamlFiles("../examples/invalid") {
			Convey(fmt.Sprintf("file %s should not validate", file), func() {
				result, err := ValidateFile(file)
				So(err, ShouldBeNil)
				So(result, ShouldNotValidate)
			})
		}
	})
}