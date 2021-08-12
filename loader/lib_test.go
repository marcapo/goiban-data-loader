package loader_test

import (
	"testing"

	data "github.com/marcapo/goiban-data"

	"github.com/marcapo/goiban-data-loader/loader"
)

func TestLoadBundesbank(t *testing.T) {
	repo := data.NewInMemoryStore()
	loader.LoadBundesbankData("../data/bundesbank.txt", repo)
	bank, err := repo.Find("DE", "79350101")

	if err != nil {
		t.Errorf("error when loading DE data: %s", err)
	}

	if bank == nil {
		t.Errorf("couldn't find known DE bank")
	}

	if bank != nil && bank.Bic != "BYLADEM1KSW" {
		t.Errorf("BIC is wrong, please check data")
	}
}

func TestLoadAustria(t *testing.T) {
	repo := data.NewInMemoryStore()
	loader.LoadAustriaData("../data/at.csv", repo)

	bank, err := repo.Find("AT", "20111")

	if err != nil {
		t.Errorf("error when loading AT data: %s", err)
	}

	if bank == nil {
		t.Errorf("couldn't find known AT bank")
	}

	if bank != nil && bank.Bic != "GIBAATWWXXX" {
		t.Errorf("AT BIC is wrong, please check data: %s", bank.Bic)
	}

}

func TestLoadSwitzerland(t *testing.T) {
	repo := data.NewInMemoryStore()
	loader.LoadSwitzerlandData("../data/ch.xlsx", repo)
}

func TestLoadLiechtenstein(t *testing.T) {
	repo := data.NewInMemoryStore()
	loader.LoadLiechtensteinData("../data/li.xlsx", repo)
}

func TestLoadLuxembourg(t *testing.T) {
	repo := data.NewInMemoryStore()
	loader.LoadLuxembourgData("../data/lu.xlsx", repo)
}

func TestLoadNetherlands(t *testing.T) {
	repo := data.NewInMemoryStore()
	loader.LoadNetherlandsData("../data/nl.xlsx", repo)
}

func TestLoadBelgium(t *testing.T) {
	repo := data.NewInMemoryStore()
	loader.LoadBelgiumData("../data/nbb.xlsx", repo)
}
