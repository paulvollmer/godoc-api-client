package godoc

import (
	"encoding/json"
	"github.com/golang/gddo/database"
	"net/http"
)

const BaseURL = "http://api.godoc.org"

type API struct {
	Client  *http.Client
	BaseURL string
}

func New() *API {
	return &API{BaseURL: BaseURL}
}

type PackagesResponse struct {
	Results []database.Package `json:"results"`
}

type ImportsResponse struct {
	Imports     []database.Package `json:"imports"`
		TestImports []database.Package `json:"testImports"`
}

func (api *API) Search(q string) (*PackagesResponse, *http.Response, error) {
	res, err := api.Client.Get(api.BaseURL + "/search?q=" + q)
	if err != nil {
		return &PackagesResponse{}, res, err
	}
	defer res.Body.Close()
	searchResponse := PackagesResponse{}
	err = json.NewDecoder(res.Body).Decode(&searchResponse)
	if err != nil {
		return &PackagesResponse{}, res, err
	}
	return &searchResponse, res, nil
}

func (api *API) Packages() (*PackagesResponse, *http.Response, error) {
	res, err := api.Client.Get(api.BaseURL + "/packages")
	if err != nil {
		return &PackagesResponse{}, res, err
	}
	defer res.Body.Close()
	searchResponse := PackagesResponse{}
	err = json.NewDecoder(res.Body).Decode(&searchResponse)
	if err != nil {
		return &PackagesResponse{}, res, err
	}
	return &searchResponse, res, nil
}

func (api *API) Importers(pkg string) (*PackagesResponse, *http.Response, error) {
	res, err := api.Client.Get(api.BaseURL + "/importers/"+pkg)
	if err != nil {
		return &PackagesResponse{}, res, err
	}
	defer res.Body.Close()
	searchResponse := PackagesResponse{}
	err = json.NewDecoder(res.Body).Decode(&searchResponse)
	if err != nil {
		return &PackagesResponse{}, res, err
	}
	return &searchResponse, res, nil
}

func (api *API) Imports(pkg string) (*ImportsResponse, *http.Response, error) {
	res, err := api.Client.Get(api.BaseURL + "/imports/"+pkg)
	if err != nil {
		return &ImportsResponse{}, res, err
	}
	defer res.Body.Close()
	importsResponse := ImportsResponse{}
	err = json.NewDecoder(res.Body).Decode(&importsResponse)
	if err != nil {
		return &ImportsResponse{}, res, err
	}
	return &importsResponse, res, nil
}
