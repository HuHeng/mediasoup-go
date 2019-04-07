package mediasoup

import "encoding/json"

type H map[string]interface{}

type Internal struct {
	RouterId      string `json:"routerId,omitempty"`
	TransportId   string `json:"transportId,omitempty"`
	ProducerId    string `json:"producerId,omitempty"`
	ConsumerId    string `json:"consumerId,omitempty"`
	RtpObserverId string `json:"rtpObserverId,omitempty"`
}

type Response struct {
	data json.RawMessage
	err  error
}

func (r Response) Result(v interface{}) error {
	if r.err != nil {
		return r.err
	}
	return json.Unmarshal([]byte(r.data), v)
}

func (r Response) Err() error {
	return r.err
}

type FetchProducerFunc func(producerId string) *Producer

type VolumeInfo struct {
	Producer *Producer
	Volume   float64
}