package lib

import "encoding/json"

const (
	// UpdateDNSRecordType represents the dns record type that can be
	// updated
	UpdateDNSRecordType = "A"
)

// UpdateDNSRecord is a type that contains the body of a dns record that
// needs to be updated
type UpdateDNSRecord struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

// NewUpdateDNSRecord creates a new update dns record
func NewUpdateDNSRecord(name, ip string) UpdateDNSRecord {
	return UpdateDNSRecord{
		Type:    UpdateDNSRecordType,
		Name:    name,
		Content: ip,
	}
}

// JSON returns the json value of this UpdateDNSRecord
func (record UpdateDNSRecord) JSON() (string, error) {
	out, err := json.Marshal(record)
	if err != nil {
		return "", err
	}
	return string(out), nil
}
