package signin

type VerifyIDTokenResponse struct {
	AuthTime int64  `json:"auth_time"`
	Issuer   string `json:"iss"`
	Audience string `json:"aud"`
	Expires  int64  `json:"exp"`
	IssuedAt int64  `json:"iat"`
	Subject  string `json:"sub,omitempty"`
	UID      string `json:"uid,omitempty"`
	Firebase struct {
		SignInProvider string                 `json:"sign_in_provider"`
		Tenant         string                 `json:"tenant"`
		Identities     map[string]interface{} `json:"identities"`
	} `json:"firebase"`
	Claims map[string]interface{} `json:"-"`
}

type SignInRequest struct {
	IDToken string `schema:"idToken"`
}

type SignInResponse struct {
	SessionCookie string
}
