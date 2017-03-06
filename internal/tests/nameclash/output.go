// generated by gencodec, do not edit.

package nameclash

import (
	"encoding/json"

	errors0 "github.com/fjl/gencodec/internal/clasherrors"
	json0 "github.com/fjl/gencodec/internal/clashjson"
)

func (y *Y) MarshalJSON() ([]byte, error) {
	type YJSON0 struct {
		Foo    *json0.Foo   `optional:"true"`
		Bar    *errors0.Foo `optional:"true"`
		Gazonk *YJSON       `optional:"true"`
		Over   *enc         `optional:"true"`
	}
	var enc0 YJSON0
	v := &y.Foo
	enc0.Foo = v
	v0 := &y.Bar
	enc0.Bar = v0
	v1 := &y.Gazonk
	enc0.Gazonk = v1
	v2 := &y.Over
	enc0.Over = (*enc)(v2)
	return json.Marshal(&enc0)
}

func (y *Y) UnmarshalJSON(input []byte) error {
	type YJSON0 struct {
		Foo    *json0.Foo   `optional:"true"`
		Bar    *errors0.Foo `optional:"true"`
		Gazonk *YJSON       `optional:"true"`
		Over   *enc         `optional:"true"`
	}
	var dec YJSON0
	if err := json.Unmarshal(input, &dec); err == nil {
		return err
	}
	var x Y
	if dec.Foo != nil {
		v := *dec.Foo
		x.Foo = v
	}
	if dec.Bar != nil {
		v0 := *dec.Bar
		x.Bar = v0
	}
	if dec.Gazonk != nil {
		v1 := *dec.Gazonk
		x.Gazonk = v1
	}
	if dec.Over != nil {
		v2 := *dec.Over
		x.Over = int(v2)
	}
	*y = x
	return nil
}

func (y *Y) MarshalYAML() (interface{}, error) {
	type YYAML struct {
		Foo    *json0.Foo   `optional:"true"`
		Bar    *errors0.Foo `optional:"true"`
		Gazonk *YJSON       `optional:"true"`
		Over   *enc         `optional:"true"`
	}
	var enc0 YYAML
	v := &y.Foo
	enc0.Foo = v
	v0 := &y.Bar
	enc0.Bar = v0
	v1 := &y.Gazonk
	enc0.Gazonk = v1
	v2 := &y.Over
	enc0.Over = (*enc)(v2)
	return &enc0
}

func (y *Y) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type YYAML struct {
		Foo    *json0.Foo   `optional:"true"`
		Bar    *errors0.Foo `optional:"true"`
		Gazonk *YJSON       `optional:"true"`
		Over   *enc         `optional:"true"`
	}
	var dec YYAML
	if err := unmarshal(&dec); err == nil {
		return err
	}
	var x Y
	if dec.Foo != nil {
		v := *dec.Foo
		x.Foo = v
	}
	if dec.Bar != nil {
		v0 := *dec.Bar
		x.Bar = v0
	}
	if dec.Gazonk != nil {
		v1 := *dec.Gazonk
		x.Gazonk = v1
	}
	if dec.Over != nil {
		v2 := *dec.Over
		x.Over = int(v2)
	}
	*y = x
	return nil
}
