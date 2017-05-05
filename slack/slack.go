/*
http://www.apache.org/licenses/LICENSE-2.0.txt

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

package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	reflect "gopkg.in/oleiade/reflections.v1"
)

const (
	Name    = "slack"
	Version = 1
)

//NewSlackPublisher returns an instance of the Slack publisher
func NewPublisher() plugin.Publisher {
	return &slackPublisher{}
}

//GetConfigPolicy returns plugins config policy
func (p *slackPublisher) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()
	policy.AddNewStringRule([]string{""}, "webhook", true)
	policy.AddNewStringRule([]string{""}, "channel", true)
	policy.AddNewStringRule([]string{""}, "username", false, plugin.SetDefaultString("webhookbot"))
	policy.AddNewStringRule([]string{""}, "publish_fields", false, plugin.SetDefaultString("Namespace|Config|Data|Description|Timestamp|Unit|Version"))

	return *policy, nil
}

// Publish publishes metric data to slack
func (p *slackPublisher) Publish(metrics []plugin.Metric, config plugin.Config) error {
	url, err := config.GetString("webhook")
	if err != nil {
		return err
	}
	publish_fields, err := config.GetString("publish_fields")
	if err != nil {
		return err
	}
	channel, err := config.GetString("channel")
	if err != nil {
		return err
	}
	username, err := config.GetString("username")
	if err != nil {
		return err
	}
	//TODO: current implementation publishes each metric separately; add publish as group (consider max message size)
	for _, metric := range metrics {
		toPublish := map[string]interface{}{}
		//TODO: validation of publish_field format to avoid errors after split string
		for _, field := range strings.Split(publish_fields, "|") {
			var value interface{}
			value, err = reflect.GetField(metric, field)
			if err != nil {
				return err
			}
			if field == "Namespace" {
				value = strings.Join(metric.Namespace.Strings(), "/")
			}
			toPublish[field] = value
		}
		message, err := json.Marshal(toPublish)
		if err != nil {
			return err
		}
		payload := map[string]string{
			"text":     string(message),
			"channel":  fmt.Sprintf("#%s", channel),
			"username": username,
		}
		js, err := json.Marshal(payload)
		if err != nil {
			return err
		}
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(js))
		if err != nil {
			return err
		}
		resp.Body.Close()
	}
	return nil
}

type slackPublisher struct{}
