package sirena

import "encoding/xml"

// GetItinReceiptsDataResponse is a Sirena response to <get_itin_receipts_data> request
type GetItinReceiptsDataResponse struct {
	Answer  GetItinReceiptsDataAnswer `xml:"answer"`
	XMLName xml.Name                  `xml:"sirena" json:"-"`
}

// GetItinReceiptsDataAnswer is an <answer> section in Sirena <get_itin_receipts_data> response
type GetItinReceiptsDataAnswer struct {
	Answer              string                            `xml:"answer,attr,omitempty"`
	GetItinReceiptsData GetItinReceiptsDataAnswerReceipts `xml:"get_itin_receipts_data>receipts"`
}

// GetItinReceiptsDataAnswerReceipts is a <receipts> element in Sirena <get_itin_receipts_data> response
type GetItinReceiptsDataAnswerReceipts struct {
	DocOfPassenger string `xml:"ticket_form>doc_of_passenger"`
	Error          *Error `xml:"error,omitempty"`
}
