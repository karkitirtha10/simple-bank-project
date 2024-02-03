package enums

type OAuthClientTypeEnum string

const (
	PERSONAL_ACCESS    OAuthClientTypeEnum = "personal access"
	CLIENT_CREDENTIALS OAuthClientTypeEnum = "client credentials"
)

//func (yo OAuthClientTypeEnum) EnumIndex() int {
//	if yo == PERSONAL_ACCESS {
//		return 1
//	}
//	return 0
//}
