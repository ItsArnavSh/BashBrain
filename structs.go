package main

type Part struct {
	Text string `json:"text"`
}

type Content struct {
	Parts []Part `json:"parts"`
	Role  string `json:"role"`
}

type Candidate struct {
	Content      Content `json:"content"`
	FinishReason string  `json:"finishReason"`
	AvgLogprobs  float64 `json:"avgLogprobs"`
}

type PromptTokenDetail struct {
	Modality   string `json:"modality"`
	TokenCount int    `json:"tokenCount"`
}

type CandidateTokenDetail struct {
	Modality   string `json:"modality"`
	TokenCount int    `json:"tokenCount"`
}

type UsageMetadata struct {
	PromptTokenCount        int                    `json:"promptTokenCount"`
	CandidatesTokenCount    int                    `json:"candidatesTokenCount"`
	TotalTokenCount         int                    `json:"totalTokenCount"`
	PromptTokensDetails     []PromptTokenDetail    `json:"promptTokensDetails"`
	CandidatesTokensDetails []CandidateTokenDetail `json:"candidatesTokensDetails"`
}

type Response struct {
	Candidates    []Candidate   `json:"candidates"`
	UsageMetadata UsageMetadata `json:"usageMetadata"`
	ModelVersion  string        `json:"modelVersion"`
}
