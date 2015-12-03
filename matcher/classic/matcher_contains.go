package classic

import (
	"github.com/gogap/spirit"
)

const (
	containsMatcherURN = "urn:spirit:matcher:label:contains"
)

var _ spirit.LabelMatcher = new(ContainsLabelMatcher)
var _ spirit.Actor = new(ContainsLabelMatcher)

type ContainsLabelMatcherConfig struct {
	Reverse bool `json:"reverse"`
}

type ContainsLabelMatcher struct {
	name string

	conf ContainsLabelMatcherConfig
}

func init() {
	spirit.RegisterLabelMatcher(containsMatcherURN, NewContainsLabelMatcher)
}

func NewContainsLabelMatcher(name string, config spirit.Map) (matcher spirit.LabelMatcher, err error) {
	conf := ContainsLabelMatcherConfig{}

	if err = config.ToObject(&conf); err != nil {
		return
	}

	matcher = &ContainsLabelMatcher{
		name: name,
		conf: conf,
	}

	return
}

func (p *ContainsLabelMatcher) Name() string {
	return p.name
}

func (p *ContainsLabelMatcher) URN() string {
	return containsMatcherURN
}

func (p *ContainsLabelMatcher) Match(la spirit.Labels, lb spirit.Labels) bool {
	if la == nil && lb == nil {
		return true
	}

	if !p.conf.Reverse {
		equalCount := 0
		for ka, va := range la {
			if vb, exist := lb[ka]; exist {
				if va == vb {
					equalCount += 1
				}
			}
		}
		if equalCount == len(lb) {
			return true
		}
	} else {
		equalCount := 0
		for kb, vb := range lb {
			if va, exist := la[kb]; exist {
				if va == vb {
					equalCount += 1
				}
			}
		}
		if equalCount == len(la) {
			return true
		}
	}

	return false
}
