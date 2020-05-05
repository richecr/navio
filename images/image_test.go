package images

import (
	"os"
	"testing"
)

var check = func(t *testing.T, expected string, result string) {
	t.Helper()
	if expected != result {
		t.Errorf("Expected %s != Result %s", expected, result)
	}
}

var clear = func() {
	RemoveDownloadedImage("alpine")
	RemoveDownloadedImage("busybox")
	RemoveDownloadedImage("ubuntu")
	os.Remove("images")
}

func TestPull(t *testing.T) {

	//note: this tests don't cover all Pull function

	t.Run("Invalid Image", func(t *testing.T) {
		imageName := "ubuntó"
		err := Pull(imageName)
		expected := "The image " + imageName + " is not available"
		result := err.Error()
		check(t, expected, result)
	})

	t.Run("A Image that already was downloaded", func(t *testing.T) {
		imageName := "alpine"
		Pull(imageName)

		err := Pull(imageName)
		expected := "The image " + imageName + " already was downloaded"
		result := err.Error()
		check(t, expected, result)
	})
	clear()
}

func TestShowDownloadedImages(t *testing.T) {
	t.Run("", func(t *testing.T) {
		if _, err := ShowDownloadedImages(); err != nil {
			t.Errorf("ERROR: on ShowDownloadedImages(): %s", err.Error())
		}
	})
}

func TestAlreadyExists(t *testing.T) {
	check := func(t *testing.T, expected bool, result bool) {
		t.Helper()
		if expected != result {
			t.Errorf("Expected %v != Result %v", expected, result)
		}
	}
	clear := func() {
		RemoveDownloadedImage("alpine")
		os.Remove("images")
	}

	Pull("alpine")
	result := AlreadyExists("alpine")
	expected := true
	check(t, expected, result)

	RemoveDownloadedImage("alpine")
	result = AlreadyExists("alpine")
	expected = false
	check(t, expected, result)

	clear()
}

func TestRemoveDownloadedImage(t *testing.T) {
	t.Run("Valid image: busybox", func(t *testing.T) {
		image := "busybox"
		if !AlreadyExists(image) {
			Pull(image)
		}

		err := RemoveDownloadedImage(image)
		if err != nil {
			t.Errorf(err.Error())
		}
		// certifies that the image was removed
		if AlreadyExists(image) {
			t.Errorf("Expected != Result ")
		}
	})
	t.Run("Invalid image: busycaixa", func(t *testing.T) {
		image := "busycaixa"
		err := RemoveDownloadedImage(image)
		if err != nil {
			t.Errorf(err.Error())
		}
	})
	t.Run("Empty image: '' ", func(t *testing.T) {
		image := ""
		err := RemoveDownloadedImage(image)
		if err.Error() != "The imageName must be a non-empty value" {
			t.Errorf(err.Error())
		}
	})
	clear()
}

func TestDescribe(t *testing.T) {
	t.Run("Unavailable Image", func(t *testing.T) {
		e := ""
		r := Describe("debianxsdsad")
		check(t, e, r)
	})
	t.Run("Alpine Image", func(t *testing.T) {
		e := "alpine\t\tv3.11\t\t2.7M"
		r := Describe("alpine")
		check(t, e, r)
	})
	t.Run("Busybox Image", func(t *testing.T) {
		e := "busybox\t\tv4.0\t\t1.5M"
		r := Describe("busybox")
		check(t, e, r)
	})
	t.Run("Ubuntu Image", func(t *testing.T) {
		e := "ubuntu\t\tv20.04\t\t90.0M"
		r := Describe("ubuntu")
		check(t, e, r)
	})
}
