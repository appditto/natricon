package model

import (
	"encoding/json"
	"testing"
)

func TestHistoryRequestSerialize(t *testing.T) {
	expected := `{"action":"account_history","account":"nano_1u96mhhk5dqy9woqhw3uajoes9j94dzcib8ohnh6ztshqy4nuby4u5iid5nm","count":6}`
	request := AccountHistoryRequest{
		BaseRequest: AccountHistoryAction,
		Account:     "nano_1u96mhhk5dqy9woqhw3uajoes9j94dzcib8ohnh6ztshqy4nuby4u5iid5nm",
		Count:       6,
	}
	jsonB, _ := json.Marshal(request)
	if string(jsonB) != expected {
		t.Errorf("Expected %s but got %s", expected, string(jsonB))
	}
}

func TestHistoryResponseDeserialize(t *testing.T) {
	toDeserialize := `{
		"account": "nano_1ipx847tk8o46pwxt5qjdbncjqcbwcc1rrmqnkztrfjy5k7z4imsrata9est",
		"history": [
		  {
			"type": "send",
			"account": "nano_38ztgpejb7yrm7rr586nenkn597s3a1sqiy3m3uyqjicht7kzuhnihdk6zpz",
			"amount": "80000000000000000000000000000000000",
			"local_timestamp": "1551532723",
			"height": "60",
			"hash": "80392607E85E73CC3E94B4126F24488EBDFEB174944B890C97E8F36D89591DC5"
		  }
		],
		"previous": "8D3AB98B301224253750D448B4BD997132400CEDD0A8432F775724F2D9821C72"
	  }`
	var resp AccountHistoryResponse
	json.Unmarshal([]byte(toDeserialize), &resp)
	if resp.Account != "nano_1ipx847tk8o46pwxt5qjdbncjqcbwcc1rrmqnkztrfjy5k7z4imsrata9est" {
		t.Errorf("Expected account %s but got %s", "nano_1ipx847tk8o46pwxt5qjdbncjqcbwcc1rrmqnkztrfjy5k7z4imsrata9est", resp.Account)
	}
	if resp.History[0].Account != "nano_38ztgpejb7yrm7rr586nenkn597s3a1sqiy3m3uyqjicht7kzuhnihdk6zpz" {
		t.Errorf("Expected account in item %s but got %s", "nano_38ztgpejb7yrm7rr586nenkn597s3a1sqiy3m3uyqjicht7kzuhnihdk6zpz", resp.History[0].Account)
	}
	if resp.History[0].Type != "send" {
		t.Errorf("Expected type send but got %s", resp.History[0].Type)
	}
	if resp.History[0].Amount != "80000000000000000000000000000000000" {
		t.Errorf("Expected amount %s but got %s", "80000000000000000000000000000000000", resp.History[0].Amount)
	}
	if resp.History[0].Hash != "80392607E85E73CC3E94B4126F24488EBDFEB174944B890C97E8F36D89591DC5" {
		t.Errorf("Expected hash %s but got %s", "80392607E85E73CC3E94B4126F24488EBDFEB174944B890C97E8F36D89591DC5", resp.History[0].Hash)
	}
}
