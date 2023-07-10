package printer_test

import (
	"testing"

	"github.com/hguerra/jacoco-check/internal/printer"
)

func TestPrint(t *testing.T) {
	t.Run("should print PASS with green color", func(t *testing.T) {
		printer.Print([]string{})
	})

	t.Run("should print violation with red color", func(t *testing.T) {
		printer.Print([]string{"a", "b", "c"})
	})
}
