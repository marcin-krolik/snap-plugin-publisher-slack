# Snap publisher plugin - Slack

This plugin supports pushing metrics to Slack via webhooks.

It's used in the [Snap framework](http://github.com/intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license)
6. [Acknowledgements](#acknowledgements)

## Getting Started

### System Requirements

* [golang 1.6+](https://golang.org/dl/) (needed only for building)

#### To build the plugin binary:
Fork https://github.com/marcin-krolik/snap-plugin-publisher-slack

Clone repo into `$GOPATH/src/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-publisher-slack.git
```

Build the plugin by running make within the cloned repo:
```
$ make
```
This builds the plugin in `./build`

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)

## Documentation

The plugin expects you to provide the following parameters:
 - `webook` eg. `https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX` more @ [Slack WebApi](https://api.slack.com/incoming-webhooks)
 - `publish_fields` eg. `"Namespace|Config|Data|Description|Timestamp|Unit|Version"`
 - `channel`
 - `username` (optional)

### Examples

TBD

### Roadmap

TBD

## Acknowledgements
* Author: [Marcin Krolik](https://github.com/marcin-krolik)

And **thank you!** Your contribution, through code and participation, is incredibly important to us.
