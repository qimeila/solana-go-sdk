package client

type GetEpochInfoResponse struct {
	AbsoluteSlot int `json:"absoluteSlot"`
	BlockHeight  int `json:"blockHeight"`
	Epoch        int `json:"epoch"`
	SlotIndex    int `json:"slotIndex"`
	SlotsInEpoch int `json:"slotsInEpoch"`
}

func (s *Client) GetEpochInfo() (GetEpochInfoResponse, error) {
	res := struct {
		GeneralResponse
		Result GetEpochInfoResponse `json:"result"`
	}{}
	s.request("getEpochInfo", []interface{}{}, &res)
	return res.Result, nil
}
