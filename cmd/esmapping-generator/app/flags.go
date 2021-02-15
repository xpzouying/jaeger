// Copyright (c) 2020 The Jaeger Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"github.com/spf13/cobra"
)

// Options represent configurable parameters for jaeger-esmapping-generator
type Options struct {
	Mapping   string
	EsVersion int64
	Shards    int64
	Replicas  int64
	EsPrefix  string
	UseILM    string // using string as util is being used in python and using bool leads to type issues.
}

const (
	mappingFlag   = "mapping"
	esVersionFlag = "es-version"
	shardsFlag    = "shards"
	replicasFlag  = "replicas"
	esPrefixFlag  = "es-prefix"
	useILMFlag    = "use-ilm"
)

// AddFlags adds flags for esmapping-generator main program
func (o *Options) AddFlags(command *cobra.Command) {
	command.Flags().StringVar(
		&o.Mapping,
		mappingFlag,
		"",
		"The index mapping the template will be applied to. Pass either jaeger-span or jaeger-service")
	command.Flags().Int64Var(
		&o.EsVersion,
		esVersionFlag,
		7,
		"The major Elasticsearch version")
	command.Flags().Int64Var(
		&o.Shards,
		shardsFlag,
		5,
		"The number of shards per index in Elasticsearch")
	command.Flags().Int64Var(
		&o.Replicas,
		replicasFlag,
		1,
		"The number of replicas per index in Elasticsearch")
	command.Flags().StringVar(
		&o.EsPrefix,
		esPrefixFlag,
		"",
		"Specifies index prefix")
	command.Flags().StringVar(
		&o.UseILM,
		useILMFlag,
		"false",
		"Set to true to use ILM for managing lifecycle of jaeger indices")

	// mark mapping flag as mandatory
	command.MarkFlagRequired(mappingFlag)
}
