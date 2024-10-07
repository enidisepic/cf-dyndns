package cloudflare

const automaticTTL = 1

type entryUpdateRequest struct {
	Content string   `json:"content"`
	Name    string   `json:"name"`
	Proxied bool     `json:"proxied"`
	Type    string   `json:"type"`
	Comment string   `json:"comment"`
	ID      string   `json:"id"`
	Tags    []string `json:"tags"`
	TTL     int      `json:"ttl"`
}

type entryUpdateResponse struct {
	Success bool `json:"success"`
}

func createEntryUpdateRequest(
	cloudflareVariables entryUpdateRequestConfig,
	ipAddress string,
) entryUpdateRequest {
	return entryUpdateRequest{
		Content: ipAddress,
		Name:    cloudflareVariables.EntryName,
		Proxied: cloudflareVariables.Proxied,
		Type:    "A",
		Comment: "",
		ID:      cloudflareVariables.ZoneID,
		Tags:    []string{},
		TTL:     automaticTTL, // Marks TTL as automatic
	}
}
