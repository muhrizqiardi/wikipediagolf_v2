package middleware

type VerifySessionCookieResponse struct {
	AuthTime int64
	Issuer   string
	Audience string
	Expires  int64
	IssuedAt int64
	Subject  string
	UID      string
	Claims   map[string]interface{}
}

type GetUserResponse struct {
	ProviderID             string `json:"providerId,omitempty"`
	UID                    string `json:"rawId,omitempty"`
	DisplayName            string `json:"displayName,omitempty"`
	Email                  string `json:"email,omitempty"`
	PhoneNumber            string `json:"phoneNumber,omitempty"`
	PhotoURL               string `json:"photoUrl,omitempty"`
	CustomClaims           map[string]interface{}
	Disabled               bool
	EmailVerified          bool
	TokensValidAfterMillis int64 // milliseconds since epoch.
	TenantID               string
}
