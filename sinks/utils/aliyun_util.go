package utils

const (
	ConfigPath         = "/var/addon/token-config"
	StsTokenTimeLayout = "2006-01-02T15:04:05Z"
)

type AKInfo struct {
	AccessKeyId     string `json:"access.key.id"`
	AccessKeySecret string `json:"access.key.secret"`
	SecurityToken   string `json:"security.token"`
	Expiration      string `json:"expiration"`
	Keyring         string `json:"keyring"`
}

func (akInfo *AKInfo) IsExpired() bool { _ = "STUB: not implemented"; return false }

func PKCS5UnPadding(origData []byte) []byte { _ = "STUB: not implemented"; return nil }

func Decrypt(s string, keyring []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func GetRegionFromEnv() (region string, err error) { _ = "STUB: not implemented"; return "", nil }

func GetOwnerAccountFromEnv() (accountId string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func ParseRegion() (string, error) { _ = "STUB: not implemented"; return "", nil }

func ParseRegionFromMeta() (string, error) { _ = "STUB: not implemented"; return "", nil }

func ParseOwnerAccountId() (string, error) { _ = "STUB: not implemented"; return "", nil }

func ParseAKInfoFromMeta() (*AKInfo, error) { _ = "STUB: not implemented"; return nil, nil }

func ParseAKInfoFromConfigPath() (*AKInfo, error) { _ = "STUB: not implemented"; return nil, nil }

//获取token config json
