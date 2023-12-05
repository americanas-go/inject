package inject

import "github.com/americanas-go/annotation"

func CollectEntries(path string) ([]annotation.Entry, error) {
	collector, err := annotation.Collect(
		annotation.WithPath(path),
		annotation.WithFilters("Inject", "Provide", "Invoke"),
	)
	if err != nil {
		return []annotation.Entry{}, err
	}

	return collector.Entries(), nil
}
