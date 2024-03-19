package arch

import (
	"go/token"
	"golang.org/x/tools/go/packages"
	"testing"
)

const loadMode = packages.NeedName |
	packages.NeedFiles |
	packages.NeedCompiledGoFiles |
	packages.NeedImports |
	packages.NeedDeps |
	packages.NeedTypes |
	packages.NeedSyntax |
	packages.NeedTypesInfo

func AssertNotDepends(t *testing.T, sourcePackage string, checkPackage string) {
	loadConfig := new(packages.Config)
	loadConfig.Mode = loadMode
	loadConfig.Fset = token.NewFileSet()
	p, err := packages.Load(loadConfig, sourcePackage)
	if err != nil {
		t.Error(err)
	}
	for pkg, _ := range p[0].Imports {
		if checkPackage == pkg {
			t.Errorf("Package %s is used on %s", checkPackage, sourcePackage)
		}
	}
}
