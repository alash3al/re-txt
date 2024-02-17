// Package text converts from text format to another
package text

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/alash3al/re-txt/handlers"
	"github.com/hashicorp/hcl"
	"gopkg.in/yaml.v2"
)

func yaml2json(c *handlers.Context) (jsonBytes []byte, err error) {
	var data interface{}

	if err := yaml.Unmarshal(c.MergeInputs(), &data); err != nil {
		return nil, err
	}

	if c.Bool("pretty") {
		jsonBytes, err = json.MarshalIndent(normalizeInterface(data), "", "   ")
	} else {
		jsonBytes, err = json.Marshal(normalizeInterface(data))
	}

	return jsonBytes, err
}

func json2yaml(c *handlers.Context) ([]byte, error) {
	var data interface{}

	if err := json.Unmarshal(c.MergeInputsAsJSON(), &data); err != nil {
		return nil, err
	}

	return yaml.Marshal(normalizeInterface(data))
}

func json2toml(c *handlers.Context) ([]byte, error) {
	var data interface{}

	if err := json.Unmarshal(c.MergeInputsAsJSON(), &data); err != nil {
		return nil, err
	}

	var writer bytes.Buffer

	if err := toml.NewEncoder(&writer).Encode(normalizeInterface(data)); err != nil {
		return nil, err
	}

	return writer.Bytes(), nil
}

func toml2json(c *handlers.Context) (jsonBytes []byte, err error) {
	var data interface{}

	if err := toml.Unmarshal(c.MergeInputs(), &data); err != nil {
		return nil, err
	}

	if c.Bool("pretty") {
		jsonBytes, err = json.MarshalIndent(normalizeInterface(data), "", "   ")
	} else {
		jsonBytes, err = json.Marshal(normalizeInterface(data))
	}

	return jsonBytes, err
}

func hcl2json(c *handlers.Context) (jsonBytes []byte, err error) {
	var data interface{}

	if err := hcl.Unmarshal(c.MergeInputs(), &data); err != nil {
		return nil, err
	}

	if c.Bool("pretty") {
		jsonBytes, err = json.MarshalIndent(normalizeInterface(data), "", "   ")
	} else {
		jsonBytes, err = json.Marshal(normalizeInterface(data))
	}

	return jsonBytes, err
}

func csv2json(c *handlers.Context) (jsonBytes []byte, err error) {
	data := []map[string]interface{}{}

	r := csv.NewReader(bytes.NewBuffer(c.MergeInputs()))
	r.Comma = rune(c.String("separator")[0])
	r.Comment = rune(c.String("comment")[0])
	r.TrimLeadingSpace = c.Bool("trim-leading-space")
	r.LazyQuotes = c.Bool("lazy-quotes")

	header := strings.Split(c.String("header"), string(r.Comma))
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	for ri, record := range records {
		if ri == 0 && c.Bool("contains-header") {
			continue
		}

		item := map[string]interface{}{}
		for i, v := range header {
			if len(record)-1 < i {
				continue
			}

			item[v] = record[i]
		}

		data = append(data, item)
	}

	if c.Bool("pretty") {
		jsonBytes, err = json.MarshalIndent(data, "", "   ")
	} else {
		jsonBytes, err = json.Marshal(data)
	}

	return jsonBytes, err
}
