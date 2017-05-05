// +build small

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
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

func TestSlackPublisher(t *testing.T) {

	Convey("Create Slack Publisher", t, func() {
		sp := NewPublisher()
		Convey("So publisher should not be nil", func() {
			So(sp, ShouldNotBeNil)
		})
		Convey("So publisher should be of InfluxPublisher type", func() {
			So(sp, ShouldHaveSameTypeAs, &slackPublisher{})
		})

		configPolicy, err := sp.GetConfigPolicy()
		Convey("sp.GetConfigPolicy() should return a config policy", func() {
			Convey("So config policy should not be nil", func() {
				So(configPolicy, ShouldNotBeNil)
			})
			Convey("So we should not get an err retreiving the config policy", func() {
				So(err, ShouldBeNil)
			})
			Convey("So config policy should be of plugin.ConfigPolicy type", func() {
				So(configPolicy, ShouldHaveSameTypeAs, plugin.ConfigPolicy{})
			})
		})
	})
}
