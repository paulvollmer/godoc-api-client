package godoc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var ApiURL = "http://api.godoc.org"

// ApiPackages godoc object
type ApiPackages struct {
	Date    time.Time
	Results []ApiPackage
}

type ApiPackage struct {
	Path string
}

func GetPackages() (*ApiPackages, []byte, error) {
	res, err := http.Get(ApiURL + "/packages")
	if err != nil {
		return &ApiPackages{}, []byte{}, err
	}
	apiJSON, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return &ApiPackages{}, []byte{}, err
	}
	var p ApiPackages
	err = json.Unmarshal(apiJSON, &p)
	p.Date = time.Now()
	return &p, apiJSON, err
}

func (a *ApiPackages) TotalPackages() int {
	return len(a.Results)
}

func (a *ApiPackages) PrettyJSON() ([]byte, error) {
	return json.MarshalIndent(a, "", "  ")
}

func ReadFile(filename string) (*ApiPackages, error) {
	d, err := ioutil.ReadFile(filename)
	if err != nil {
		return &ApiPackages{}, err
	}
	var p ApiPackages
	err = json.Unmarshal(d, &p)
	return &p, nil
}

func (a *ApiPackages) WriteFile(s string) error {
	d, err := a.PrettyJSON()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(s, d, 0755)
	return err
}

func (a *ApiPackages) PrettyPrint() {
	fmt.Println("URL ", ApiURL)
	fmt.Println("DATE", a.Date)
	fmt.Printf("total packages: %v\n", a.TotalPackages())
}
