/*
Copyright 2021 kubeflow.org.

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

package logger

import (
	"encoding/json"
	"fmt"
)

type KafkaConnectType struct {
	Schema  KafkaConnectSchema  `json:"schema"`
	Payload KafkaConnectPayload `json:"payload"`
}

type KafkaConnectSchema struct {
	Type     string                     `json:"type"`
	Fields   []KafkaConnectSchemaFields `json:"fields"`
	Optional bool                       `json:"optional"`
	Name     string                     `json:"name"`
}

type KafkaConnectSchemaFields struct {
	Type     string `json:"type"`
	Field    string `json:"field"`
	Optional bool   `json:"optional"`
}

type KafkaConnectPayload struct {
	Body          string `json:"body"`
	TransactionID string `json:"transaction_id"`
	ModelName     string `json:"model_name"`
}

// CreateKafkaConnectPayload converts regular payload to Kafka Connect Payload
func CreateKafkaConnectPayload(logReq LogRequest) ([]byte, error) {
	payloadBody := string(*logReq.Bytes)
	kafkaConnectBody := &KafkaConnectType{
		Schema: KafkaConnectSchema{
			Type: "struct",
			Fields: []KafkaConnectSchemaFields{
				{
					Type:     "string",
					Field:    "body",
					Optional: false,
				},
				{
					Type:     "string",
					Field:    "transaction_id",
					Optional: false,
				},
				{
					Type:     "string",
					Field:    "model_name",
					Optional: false,
				},
			},
			Optional: false,
			Name:     "ksql.users",
		},
		Payload: KafkaConnectPayload{
			Body:          payloadBody,
			TransactionID: logReq.Id,
			ModelName:     logReq.InferenceService,
		},
	}
	payload, err := json.Marshal(kafkaConnectBody)
	if err != nil {
		return nil, fmt.Errorf("while creating Kafka Connect payload: %s", err)
	}
	return payload, nil
}
