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
* Set up the [Snap framework](https://github.com/intelsdi-x/snap#getting-started)

## Documentation

The plugin expects you to provide the following parameters:
 - `webhook` eg. `https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX` more @ [Slack WebApi](https://api.slack.com/incoming-webhooks)
 - `publish_fields` eg. `"Namespace|Config|Data|Description|Timestamp|Unit|Version"` (optional)
 - `channel`
 - `username` (optional)

### Examples

Example of using snap-plugin-publisher-slack to store metrics collected by snap-plugin-colector-psutil.
Make sure you have access to Slack and [set up incoming webhook integration](https://api.slack.com/incoming-webhooks)

Ensure [Snap daemon is running](https://github.com/intelsdi-x/snap#running-snap):
```
$ snapteld -l -t 0 &
```
Download and load Snap plugins:
```
$ snaptel plugin load snap-plugin-collector-psutil
$ snaptel plugin load snap-plugin-publisher-slack
```
Create a task:
```
$ snaptel task create -t psutil-slack.yml
Using task manifest to create task
Task created
ID: 3b2bc29c-8256-4ca0-a958-05f5d41f6f11
Name: Task-3b2bc29c-8256-4ca0-a958-05f5d41f6f11
State: Running
```
You may view example tasks [here](examples/tasks/psutil-slack.yml).

Ensure the task is running and collecting metrics:
```
$ snaptel task watch 3b2bc29c-8256-4ca0-a958-05f5d41f6f11
```

### Roadmap
There isn't a current roadmap for this plugin. As we launch this plugin, we do not have any outstanding requirements for the next release. 
If you have a feature request, please add it as an [issue](https://github.com/marcin-krolik/snap-plugin-publisher-slack/issues/new) and/or submit a [pull request](https://github.com/marcin-krolik/snap-plugin-publisher-slack/pulls).

## Acknowledgements
* Author: [Marcin Krolik](https://github.com/marcin-krolik)