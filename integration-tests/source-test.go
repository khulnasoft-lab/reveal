package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/khulnasoft-labs/reveal/testutils"
	folderutil "github.com/khulnasoft-labs/utils/folder"
)

var (
	ConfigFile = filepath.Join(folderutil.HomeDirOrDefault("."), ".config/reveal/provider-config.yaml")
)

type censysTestcases struct{}

func (h censysTestcases) Execute() error {
	token := os.Getenv("CENSYS_API_KEY")
	if token == "" {
		return errors.New("missing censys api key")
	}
	censysToken := fmt.Sprintf(`censys: [%s]`, token)
	_ = os.WriteFile(ConfigFile, []byte(censysToken), 0644)
	defer os.RemoveAll(ConfigFile)
	results, err := testutils.RunRevealAndGetResults(debug, "-censys", "'services.software.vendor=Grafana'")
	if err != nil {
		return err
	}
	return expectResultsGreaterThanCount(results, 0)
}

type shodanTestcases struct{}

func (h shodanTestcases) Execute() error {
	token := os.Getenv("SHODAN_API_KEY")
	if token == "" {
		return errors.New("missing shodan api key")
	}
	shodanToken := fmt.Sprintf(`shodan: [%s]`, token)
	_ = os.WriteFile(ConfigFile, []byte(shodanToken), 0644)
	defer os.RemoveAll(ConfigFile)
	results, err := testutils.RunRevealAndGetResults(debug, "-shodan", "'title:\"Grafana\"'")
	if err != nil {
		return err
	}
	err = expectResultsGreaterThanCount(results, 0)
	if err != nil {
		return err
	}
	results, err = testutils.RunRevealAndGetResults(debug, "-shodan", "'org:\"Something, Inc.\"'")
	if err != nil {
		return err
	}
	return expectResultsGreaterThanCount(results, 1)
}

type zoomeyeTestcases struct{}

func (h zoomeyeTestcases) Execute() error {
	token := os.Getenv("ZOOMEYE_API_KEY")
	if token == "" {
		return errors.New("missing zoomeye api key")
	}
	zoomeyeToken := fmt.Sprintf(`zoomeye: [%s]`, token)
	_ = os.WriteFile(ConfigFile, []byte(zoomeyeToken), 0644)
	defer os.RemoveAll(ConfigFile)
	results, err := testutils.RunRevealAndGetResults(debug, "-zoomeye", "'app:\"Atlassian JIRA\"'")
	if err != nil {
		return err
	}
	return expectResultsGreaterThanCount(results, 0)
}

type fofaTestcases struct{}

func (h fofaTestcases) Execute() error {
	token := os.Getenv("FOFA_API_KEY")
	if token == "" {
		return errors.New("missing fofa api key")
	}
	fofaToken := fmt.Sprintf(`fofa: [%s]`, token)
	_ = os.WriteFile(ConfigFile, []byte(fofaToken), 0644)
	defer os.RemoveAll(ConfigFile)
	results, err := testutils.RunRevealAndGetResults(debug, "-fofa", "'app=Grafana'")
	if err != nil {
		return err
	}
	return expectResultsGreaterThanCount(results, 0)
}

// type hunterTestcases struct{}

// func (h hunterTestcases) Execute() error {
// 	results, err := testutils.RunRevealAndGetResults(debug, "-hunter", "'Grafana'")
// 	if err != nil {
// 		return err
// 	}
// 	return expectResultsGreaterOrEqualToCount(results, 0)
// }

type quakeTestcases struct{}

func (h quakeTestcases) Execute() error {
	token := os.Getenv("QUAKE_API_KEY")
	if token == "" {
		return errors.New("missing quake api key")
	}
	quakeToken := fmt.Sprintf(`quake: [%s]`, token)
	_ = os.WriteFile(ConfigFile, []byte(quakeToken), 0644)
	defer os.RemoveAll(ConfigFile)
	results, err := testutils.RunRevealAndGetResults(debug, "-quake", "'Grafana'")
	if err != nil {
		return err
	}
	return expectResultsGreaterThanCount(results, 0)
}

type netlasTestcases struct{}

func (h netlasTestcases) Execute() error {
	token := os.Getenv("NETLAS_API_KEY")
	if token == "" {
		return errors.New("missing netlas api key")
	}
	netlasToken := fmt.Sprintf(`netlas: [%s]`, token)
	_ = os.WriteFile(ConfigFile, []byte(netlasToken), 0644)
	defer os.RemoveAll(ConfigFile)
	results, err := testutils.RunRevealAndGetResults(debug, "-netlas", "'Grafana'")
	if err != nil {
		return err
	}
	return expectResultsGreaterThanCount(results, 0)
}

type criminalipTestcases struct{}

func (h criminalipTestcases) Execute() error {
        token := os.Getenv("CRIMINALIP_API_KEY")
        if token == "" {
                return errors.New("missing criminalip api key")
        }
        criminalipToken := fmt.Sprintf(`criminalip: [%s]`, token)
        _ = os.WriteFile(ConfigFile, []byte(criminalipToken), 0644)
        defer os.RemoveAll(ConfigFile)
        results, err := testutils.RunRevealAndGetResults(debug, "-criminalip", "'Grafana'")
        if err != nil {
                return err
        }
        return expectResultsGreaterThanCount(results, 0)
}
