package proto

import (
	"encoding/json"

	"github.com/pion/webrtc/v2"
)

type ClientUserInfo struct {
	Name string `json:"name"`
}

func (m *ClientUserInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (m *ClientUserInfo) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

type RoomInfo struct {
	RID RID `json:"rid"`
	UID UID `json:"uid"`
}

type RTCInfo struct {
	Description webrtc.SessionDescription `json:"description"`
}

type PublishOptions struct {
	Codec       string `json:"codec"`
	Resolution  string `json:"resolution"`
	Bandwidth   int    `json:"bandwidth"`
	Audio       bool   `json:"audio"`
	Video       bool   `json:"video"`
	Screen      bool   `json:"screen"`
	TransportCC bool   `json:"transportCC,omitempty"`
}

type SubscribeOptions struct {
	Bandwidth   int  `json:"bandwidth"`
	TransportCC bool `json:"transportCC"`
}

type Stream struct {
	ID     string      `json:"id"`
	Tracks []TrackInfo `json:"tracks"`
}

/// Messages ///

type JoinMsg struct {
	RoomInfo
	Info ClientUserInfo `json:"info"`
}

type LeaveMsg struct {
	RoomInfo
	Info ClientUserInfo `json:"info"`
}

type PublishMsg struct {
	RoomInfo
	RTCInfo
	Options PublishOptions `json:"options"`
}

type PublishResponseMsg struct {
	MediaInfo
	RTCInfo
	Stream Stream `json:"stream"`
}

type UnpublishMsg struct {
	MediaInfo
}

type SubscribeMsg struct {
	MediaInfo
	RTCInfo
	Options SubscribeOptions
}

type SubscribeResponseMsg struct {
	MediaInfo
	RTCInfo
}

type UnsubscribeMsg struct {
	MediaInfo
}

type BroadcastMsg struct {
	RoomInfo
	Info json.RawMessage `json:"info"`
}

type TrickleMsg struct {
	MediaInfo
	Info    json.RawMessage `json:"info"`
	Trickle json.RawMessage `json:"trickle"`
}

type StreamAddMsg struct {
	MediaInfo
	Info   ClientUserInfo `json:"info"`
	Stream Stream         `json:"stream"`
}

type StreamRemoveMsg struct {
	MediaInfo
}

type GetMediaInfoResponseMsg struct {
	MediaInfo
	Stream Stream `json:"stream"`
}
