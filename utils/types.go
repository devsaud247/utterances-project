package utils

type Utterance struct {
	Speaker     string `json:"speaker"`
	Text        string `json:"text"`
	TimestampMs int64  `json:"timestampMs"`
	Partial     bool   `json:"isPartial"`
}
