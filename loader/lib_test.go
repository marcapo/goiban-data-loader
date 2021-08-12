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

func TestLoadBundesbankFromDefaultPath(t *testing.T) {
	repo := data.NewInMemoryStore()
	loader.LoadBundesbankData(loader.DefaultBundesbankPath(), repo)
}

func TestLoadAustriaFromDefaultPath(t *testing.T) {
	repo := data.NewInMemoryStore()
	loader.LoadAustriaData(loader.DefaultAustriaPath(), repo)

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

func TestLoadSwitzerlandFromDefaultPath(t *testing.T) {
	repo := data.NewInMemoryStore()
	loader.LoadSwitzerlandData(loader.DefaultSwitzerlandPath(), repo)
}

func TestLoadLiechtensteinFromDefaultPath(t *testing.T) {
	repo := data.NewInMemoryStore()
	loader.LoadLiechtensteinData(loader.DefaultLiechtensteinPath(), repo)
}

func TestLoadLuxembourgFromDefaultPath(t *testing.T) {
	repo := data.NewInMemoryStore()
	loader.LoadLuxembourgData(loader.DefaultLuxembourgPath(), repo)
}

func TestLoadNetherlandsFromDefaultPath(t *testing.T) {
	repo := data.NewInMemoryStore()
	loader.LoadNetherlandsData(loader.DefaultNetherlandsPath(), repo)
}

func TestLoadBelgiumFromDefaultPath(t *testing.T) {
	repo := data.NewInMemoryStore()
	loader.LoadBelgiumData(loader.DefaultBelgiumPath(), repo)
}
