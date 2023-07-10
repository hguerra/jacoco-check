package validator_test

import (
	"testing"

	"github.com/hguerra/jacoco-check/internal/validator"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("should return an error if xmlReportPath is empty", func(t *testing.T) {
		res, err := validator.Validate("", []string{}, 0.0, 0.0)
		assert.ErrorContains(t, err, "xmlReportPath is mandatory")
		assert.Empty(t, res)
	})

	t.Run("should return an error if filesChanged is empty", func(t *testing.T) {
		res, err := validator.Validate("jacoco.xml", []string{}, 0.0, 0.0)
		assert.ErrorContains(t, err, "filesChanged should not be empty")
		assert.Empty(t, res)
	})

	t.Run("should return an error if xmlReportPath does not exists", func(t *testing.T) {
		res, err := validator.Validate("jacoco.xml", []string{"abc.txt"}, 0.0, 0.0)
		assert.ErrorContains(t, err, "open jacoco.xml: no such file or directory")
		assert.Empty(t, res)
	})

	t.Run("should return an error for invalid xml content", func(t *testing.T) {
		res, err := validator.Validate("../../test/data/jacoco-invalid.xml", []string{"abc.txt"}, 0.0, 0.0)
		assert.ErrorContains(t, err, "EOF")
		assert.Empty(t, res)
	})

	t.Run("should print a table with validation result", func(t *testing.T) {
		res, err := validator.Validate(
			"../../test/data/jacoco.xml",
			[]string{".editorconfig", "src/test/java/br/com/company/crm/infra/validator/ObjectValidatorTest.kt"},
			0.0,
			0.0,
		)
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("should return an error if code coverage on overall code less than x%", func(t *testing.T) {
		res, err := validator.Validate(
			"../../test/data/jacoco.xml",
			[]string{".editorconfig", "src/test/java/br/com/company/crm/infra/validator/ObjectValidatorTest.kt"},
			0.95,
			0.0,
		)
		assert.ErrorContains(t, err, "Code coverage on overall code 0%, required 95%")
		assert.Empty(t, res)
	})

	t.Run("should return an error if code coverage on new code less than x%", func(t *testing.T) {
		res, err := validator.Validate(
			"../../test/data/jacoco.xml",
			[]string{".editorconfig", "src/test/java/br/com/company/crm/infra/validator/ObjectValidatorTest.kt"},
			0.0,
			0.95,
		)
		assert.ErrorContains(t, err, "Code coverage on new code 0%, required 95%")
		assert.Empty(t, res)
	})
}
