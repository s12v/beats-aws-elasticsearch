package main

import (
	"github.com/elastic/beats/libbeat/outputs"
	"github.com/elastic/beats/libbeat/plugin"

	"github.com/s12v/awsbeats/elasticsearch"
)

var Bundle = plugin.Bundle(
	outputs.Plugin("aws_elasticsearch", elasticsearch.New),
)
