package httperrors

type ErrorInfo struct {
	Reason   string            `json:"reason,omitempty"`
	Domain   string            `json:"domain,omitempty"`
	Metadata map[string]string `json:"message,omitempty"`
}
