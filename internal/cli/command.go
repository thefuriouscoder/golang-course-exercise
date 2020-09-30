package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/mohae/struct2csv"
	"github.com/spf13/cobra"
	"github.com/thefuriouscoder/golang-exercise/internal/model"
	"github.com/tidwall/pretty"
)

const idFlag = "id"
const minIbuFlag = "min-ibu"
const maxIbuFlag = "max-ibu"
const minAbvFlag = "min-abv"
const maxAbvFlag = "max-abv"
const yeastFlag = "yeast"
const hopsFlag = "hops"
const maltsFlag = "malts"
const nameFlag = "name"
const formatFlag = "format"
const outputFlag = "output"

// CobraFn function definition of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

// InitBeerCmd initialize beers command
func InitBeerCmd(repository model.PunkRepo) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "beers",
		Short: "Print data about Punk Brewery beers. Use id flag to retrieve info for a given beer.",
		Run:   runBeersFn(repository),
	}

	cmd.Flags().StringP(idFlag, "i", "", "id of the beer")
	cmd.Flags().StringP(formatFlag, "f", "json", "format to export data. Available formats: json csv")
	cmd.Flags().StringP(outputFlag, "o", "", "path to file for exporting data")

	return cmd
}

// InitSearchCmd initialize search command
func InitSearchCmd(repository model.PunkRepo) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search",
		Short: "Seacrh Punk Brewery beers.",
		Run:   runSearchFn(repository),
	}

	cmd.Flags().StringP(nameFlag, "n", "", "name of the beer")
	cmd.Flags().StringP(hopsFlag, "", "", "hops contained in the beer")
	cmd.Flags().StringP(maltsFlag, "m", "", "malts contained in the beer")
	cmd.Flags().StringP(yeastFlag, "y", "", "yeast contained in the beer")
	cmd.Flags().StringP(minIbuFlag, "", "", "min ibu of the beer")
	cmd.Flags().StringP(maxIbuFlag, "", "", "max ibu of the beer")
	cmd.Flags().StringP(minAbvFlag, "", "", "min abv of the beer")
	cmd.Flags().StringP(maxAbvFlag, "", "", "max abv of the beer")
	cmd.Flags().StringP(formatFlag, "f", "json", "format to export data. Available formats: json csv")
	cmd.Flags().StringP(outputFlag, "o", "", "path to file for exporting data")

	return cmd
}

func runBeersFn(repository model.PunkRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(idFlag)
		format, _ := cmd.Flags().GetString(formatFlag)
		output, _ := cmd.Flags().GetString(outputFlag)

		if id != "" {
			id, _ := strconv.Atoi(id)
			beers, _ := repository.GetBeer(id)
			if output != "" {
				saveResponse(beers, format, output)
			} else {
				printResponse(beers, format)
			}
		} else {
			beers, _ := repository.GetBeers()
			if output != "" {
				saveResponse(beers, format, output)
			} else {
				printResponse(beers, format)
			}
		}

		return
	}
}

func runSearchFn(repository model.PunkRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString(nameFlag)
		yeast, _ := cmd.Flags().GetString(yeastFlag)
		malts, _ := cmd.Flags().GetString(maltsFlag)
		hops, _ := cmd.Flags().GetString(hopsFlag)
		minIbu, _ := cmd.Flags().GetString(minIbuFlag)
		maxIbu, _ := cmd.Flags().GetString(maxIbuFlag)
		minAbv, _ := cmd.Flags().GetString(minAbvFlag)
		maxAbv, _ := cmd.Flags().GetString(maxAbvFlag)

		format, _ := cmd.Flags().GetString(formatFlag)
		output, _ := cmd.Flags().GetString(outputFlag)

		var searchTerms = make(map[string]string)

		if name != "" {
			searchTerms["beer_name"] = name
		}

		if yeast != "" {
			searchTerms["yeast"] = yeast
		}

		if malts != "" {
			searchTerms["malt"] = malts
		}

		if hops != "" {
			searchTerms["hops"] = hops
		}

		if minIbu != "" {
			searchTerms["ibu_gt"] = minIbu
		}

		if maxIbu != "" {
			searchTerms["ibu_lt"] = maxIbu
		}

		if minAbv != "" {
			searchTerms["abv_gt"] = minAbv
		}

		if maxAbv != "" {
			searchTerms["abv_lt"] = maxAbv
		}

		beers, _ := repository.Search(searchTerms)
		if output != "" {
			saveResponse(beers, format, output)
		} else {
			printResponse(beers, format)
		}

		return
	}
}

func printResponse(beers []model.Beer, format string) {
	if format == "csv" {
		printCSV(beers)
	} else {
		printJSON(beers)
	}

}

func saveResponse(beers []model.Beer, format, path string) {
	if format == "csv" {
		saveCSV(beers, path)
	} else {
		saveJSON(beers, path)
	}
}

func saveJSON(beers []model.Beer, path string) {
	content, _ := json.MarshalIndent(beers, "", " ")
	_ = ioutil.WriteFile(path, content, 0644)
}

func saveCSV(beers []model.Beer, path string) {
	buff := &bytes.Buffer{}
	writer := struct2csv.NewWriter(buff)
	writer.WriteColNames(beers)
	writer.WriteStructs(beers)
	writer.Flush()
	_ = ioutil.WriteFile(path, buff.Bytes(), 0644)
}

func printJSON(beers []model.Beer) {
	json, _ := json.Marshal(beers)
	colorized := pretty.Color(pretty.Pretty(json), nil)
	fmt.Println(string(colorized))
}

func printCSV(beers []model.Beer) {
	buff := &bytes.Buffer{}
	writer := struct2csv.NewWriter(buff)
	writer.WriteColNames(beers)
	writer.WriteStructs(beers)
	writer.Flush()
	fmt.Println(buff.String())
}
