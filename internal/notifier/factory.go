/*
Copyright 2020 The Flux CD contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package notifier

import (
	"fmt"

	"github.com/fluxcd/notification-controller/api/v1alpha1"
)

type Factory struct {
	URL      string
	Username string
	Channel  string
	Token    string
}

func NewFactory(url string, username string, channel string, token string) *Factory {
	return &Factory{
		URL:      url,
		Channel:  channel,
		Username: username,
		Token:    token,
	}
}

func (f Factory) Notifier(provider string) (Interface, error) {
	if f.URL == "" {
		return &NopNotifier{}, nil
	}

	var n Interface
	var err error
	switch provider {
	case v1alpha1.GenericProvider:
		n, err = NewForwarder(f.URL)
	case v1alpha1.SlackProvider:
		n, err = NewSlack(f.URL, f.Username, f.Channel)
	case v1alpha1.DiscordProvider:
		n, err = NewDiscord(f.URL, f.Username, f.Channel)
	case v1alpha1.RocketProvider:
		n, err = NewRocket(f.URL, f.Username, f.Channel)
	case v1alpha1.MSTeamsProvider:
		n, err = NewMSTeams(f.URL)
	case v1alpha1.GitHubProvider:
		n, err = NewGitHub(f.URL, f.Token)
	default:
		err = fmt.Errorf("provider %s not supported", provider)
	}

	if err != nil {
		n = &NopNotifier{}
	}
	return n, err
}
