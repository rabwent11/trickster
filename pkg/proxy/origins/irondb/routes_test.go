/*
 * Copyright 2018 Comcast Cable Communications Management, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package irondb

import (
	"testing"

	"github.com/tricksterproxy/trickster/pkg/proxy/request"
	tu "github.com/tricksterproxy/trickster/pkg/util/testing"
)

func TestRegisterHandlers(t *testing.T) {
	c := &Client{}
	c.registerHandlers()
	if _, ok := c.handlers[mnCAQL]; !ok {
		t.Errorf("expected to find handler named: %s", mnCAQL)
	}
}

func TestHandlers(t *testing.T) {
	c := &Client{}
	m := c.Handlers()
	if _, ok := m[mnCAQL]; !ok {
		t.Errorf("expected to find handler named: %s", mnCAQL)
	}
}

func TestDefaultPathConfigs(t *testing.T) {

	client := &Client{name: "test"}
	ts, _, r, hc, err := tu.NewTestInstance("", client.DefaultPathConfigs, 200, "{}", nil, "irondb", "/health", "debug")
	rsc := request.GetResources(r)
	rsc.OriginClient = client
	client.config = rsc.OriginConfig
	client.webClient = hc
	defer ts.Close()
	if err != nil {
		t.Error(err)
	}

	if _, ok := client.config.Paths["/"]; !ok {
		t.Errorf("expected to find path named: %s", "/")
	}

	const expectedLen = 10
	if len(client.config.Paths) != expectedLen {
		t.Errorf("expected ordered length to be: %d got %d", expectedLen, len(client.config.Paths))
	}

}
