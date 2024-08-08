package yougile_api_wrapp

import (
	"net/url"
)

type KeyWordArguments map[string]string

func DefaultsKwargs() KeyWordArguments {
	return make(KeyWordArguments)
}

func (kwargs KeyWordArguments) ToURLValues() url.Values {
	v := url.Values{}
	for key, value := range kwargs {
		v.Set(key, value)
	}
	return v
}

func flattenArguments(extraArgs []KeyWordArguments) (kwargs KeyWordArguments) {
	kwargs = make(KeyWordArguments)
	kwargs.flatten(extraArgs)
	return
}

func (kwargs KeyWordArguments) flatten(extraKwarg []KeyWordArguments) {
	for _, extraKwarg := range extraKwarg {
		for key, val := range extraKwarg {
			kwargs[key] = val
		}
	}
}
