package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/maas/gomaasclient/entity"
)

type Space struct {
	ApiClient ApiClient
}

func (s *Space) client(id int) ApiClient {
	return s.ApiClient.GetSubObject("spaces").GetSubObject(fmt.Sprintf("%v", id))
}

func (s *Space) Get(id int) (space *entity.Space, err error) {
	space = new(entity.Space)
	err = s.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, space)
	})
	return
}

func (s *Space) Update(id int, name string) (space *entity.Space, err error) {
	space = new(entity.Space)
	qsp := url.Values{}
	qsp.Set("name", name)
	err = s.client(id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, space)
	})
	return
}

func (s *Space) Delete(id int) error {
	return s.client(id).Delete()
}
