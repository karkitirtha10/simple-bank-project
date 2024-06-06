package enums

type OAuthClientTypeShortNameEnum string

const (
	PERSONAL OAuthClientTypeShortNameEnum = "personal"
	CLIENT   OAuthClientTypeShortNameEnum = "client"
)

func (yo OAuthClientTypeShortNameEnum) ToOAuthClientTypeEnum() OAuthClientTypeEnum {
	return map[OAuthClientTypeShortNameEnum]OAuthClientTypeEnum{
		PERSONAL: PERSONAL_ACCESS,
		CLIENT:   CLIENT_CREDENTIALS,
	}[yo]
}
