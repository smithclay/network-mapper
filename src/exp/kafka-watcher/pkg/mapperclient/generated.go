// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package mapperclient

import (
	"context"
	"time"

	"github.com/Khan/genqlient/graphql"
)

type KafkaMapperResult struct {
	SrcIp           string    `json:"srcIp"`
	ServerPodName   string    `json:"serverPodName"`
	ServerNamespace string    `json:"serverNamespace"`
	Topic           string    `json:"topic"`
	Operation       string    `json:"operation"`
	LastSeen        time.Time `json:"lastSeen"`
}

// GetSrcIp returns KafkaMapperResult.SrcIp, and is useful for accessing the field via an interface.
func (v *KafkaMapperResult) GetSrcIp() string { return v.SrcIp }

// GetServerPodName returns KafkaMapperResult.ServerPodName, and is useful for accessing the field via an interface.
func (v *KafkaMapperResult) GetServerPodName() string { return v.ServerPodName }

// GetServerNamespace returns KafkaMapperResult.ServerNamespace, and is useful for accessing the field via an interface.
func (v *KafkaMapperResult) GetServerNamespace() string { return v.ServerNamespace }

// GetTopic returns KafkaMapperResult.Topic, and is useful for accessing the field via an interface.
func (v *KafkaMapperResult) GetTopic() string { return v.Topic }

// GetOperation returns KafkaMapperResult.Operation, and is useful for accessing the field via an interface.
func (v *KafkaMapperResult) GetOperation() string { return v.Operation }

// GetLastSeen returns KafkaMapperResult.LastSeen, and is useful for accessing the field via an interface.
func (v *KafkaMapperResult) GetLastSeen() time.Time { return v.LastSeen }

type KafkaMapperResults struct {
	Results []KafkaMapperResult `json:"results"`
}

// GetResults returns KafkaMapperResults.Results, and is useful for accessing the field via an interface.
func (v *KafkaMapperResults) GetResults() []KafkaMapperResult { return v.Results }

// __reportKafkaMapperResultsInput is used internally by genqlient
type __reportKafkaMapperResultsInput struct {
	Results KafkaMapperResults `json:"results"`
}

// GetResults returns __reportKafkaMapperResultsInput.Results, and is useful for accessing the field via an interface.
func (v *__reportKafkaMapperResultsInput) GetResults() KafkaMapperResults { return v.Results }

// reportKafkaMapperResultsResponse is returned by reportKafkaMapperResults on success.
type reportKafkaMapperResultsResponse struct {
	ReportKafkaMapperResults bool `json:"reportKafkaMapperResults"`
}

// GetReportKafkaMapperResults returns reportKafkaMapperResultsResponse.ReportKafkaMapperResults, and is useful for accessing the field via an interface.
func (v *reportKafkaMapperResultsResponse) GetReportKafkaMapperResults() bool {
	return v.ReportKafkaMapperResults
}

func reportKafkaMapperResults(
	ctx context.Context,
	client graphql.Client,
	results KafkaMapperResults,
) (*reportKafkaMapperResultsResponse, error) {
	req := &graphql.Request{
		OpName: "reportKafkaMapperResults",
		Query: `
mutation reportKafkaMapperResults ($results: KafkaMapperResults!) {
	reportKafkaMapperResults(results: $results)
}
`,
		Variables: &__reportKafkaMapperResultsInput{
			Results: results,
		},
	}
	var err error

	var data reportKafkaMapperResultsResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
