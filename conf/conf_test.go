package conf

import (
	"testing"
)

func TestDuplicates(t *testing.T) {
	newImport := func(pkg, version, repo string) Import {
		return Import{Package: pkg, Version: version, Repo: repo}
	}

	testData := []struct {
		imports    []Import
		duplicates int
	}{
		{[]Import{
			newImport("package1", "version1", ""),
		}, 0},
		{[]Import{
			newImport("package1", "version1", ""),
			newImport("package2", "version1", "repoA"),
		}, 0},
		{[]Import{
			newImport("package1", "version1", ""),
			newImport("package2", "version1", "repoA"),
			newImport("package1", "version1", ""),
		}, 1},
		{[]Import{
			newImport("package1", "version1", ""),
			newImport("package2", "version1", "repoA"),
			newImport("package1", "version1", ""),
			newImport("package1", "version1", ""),
		}, 2},
		{[]Import{
			newImport("package1", "version1", ""),
			newImport("package2", "version1", "repoA"),
			newImport("package1", "version1", ""),
			newImport("package1", "version1", ""),
			newImport("package2", "version2", "repoB"),
			newImport("package3", "version1", "repoA"),
		}, 3},
	}

	for i, d := range testData {
		trash := Conf{
			Package:   "",
			Imports:   d.imports,
			Excludes:  nil,
			Packages:  nil,
			ImportMap: nil,
			confFile:  "",
			yamlType:  false,
		}
		trash.Dedupe()

		if d.duplicates != len(d.imports)-len(trash.Imports) {
			t.Errorf("Case %d failed: expected %d duplicates but removed %d", i, d.duplicates, len(d.imports)-len(trash.Imports))
		}

	}

}
