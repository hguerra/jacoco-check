/*
Copyright © 2023 Heitor Carneiro <heitorgcarneiro@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"

	"github.com/hguerra/jacoco-check/internal/validator"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		xmlReportPath, err := cmd.Flags().GetString("xml")
		if err != nil {
			log.Fatal(err)
		}

		filesChanged, err := cmd.Flags().GetStringSlice("files")
		if err != nil {
			log.Fatal(err)
		}

		coverageOverallCode, err := cmd.Flags().GetFloat32("coverage-overall-code")
		if err != nil {
			log.Fatal(err)
		}

		coverageNewCode, err := cmd.Flags().GetFloat32("coverage-new-code")
		if err != nil {
			log.Fatal(err)
		}

		res, err := validator.Validate(xmlReportPath, filesChanged, coverageOverallCode, coverageNewCode)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	var minCodeCoverageOverall float32
	var minCodeCoverageNewCode float32 = 0.8

	checkCmd.Flags().StringP(
		"xml",
		"x",
		"",
		"Jacoco XML report path",
	)

	checkCmd.Flags().StringSliceP(
		"files",
		"f",
		[]string{},
		"Files changed in a pull request",
	)

	checkCmd.Flags().Float32P(
		"coverage-overall-code",
		"o",
		minCodeCoverageOverall,
		"Code coverage on overall code greater than 0% (where 0.0 represents 0%)",
	)

	checkCmd.Flags().Float32P(
		"coverage-new-code",
		"n",
		minCodeCoverageNewCode,
		"Code coverage on new code greater than 80% (where 0.8 represents 80%)",
	)
}
