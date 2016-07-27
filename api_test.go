package godoc

import (
	"testing"
)

func Test_Api(t *testing.T) {
	packages, _, err := GetPackages()
	if err != nil {
		t.Error(err)
	}

	if packages.TotalPackages() == 0 {
		t.Error("Something went wrong")
	}

	err = packages.WriteFile(packages.Date.String() + "_godoc_api_packages.json")
	if err != nil {
		t.Error(err)
	}

}
