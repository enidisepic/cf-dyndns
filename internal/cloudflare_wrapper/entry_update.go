package cloudflare_wrapper

type entryUpdateRequest struct {
	Content string   `json:"content"`
	Name    string   `json:"name"`
	Proxied bool     `json:"proxied"`
	Type    string   `json:"type"`
	Comment string   `json:"comment"`
	Id      string   `json:"id"`
	Tags    []string `json:"tags"`
	Ttl     int      `json:"ttl"`
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
		Proxied: false,
		Type:    "A",
		Comment: "",
		Id:      cloudflareVariables.ZoneId,
		Tags:    []string{},
		Ttl:     1, // Marks TTL as automatic
	}
}
