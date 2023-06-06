package constants

type EnvironmentVariable int32

const (
	EnvAccessToken EnvironmentVariable = iota
	EnvUserCountryCode
	EnvUserID
)

func (e EnvironmentVariable) Path() string {
	pathMap := map[EnvironmentVariable]string{
		EnvAccessToken:     "access_token.txt",
		EnvUserCountryCode: "user_country_code.txt",
		EnvUserID:          "user_id.txt",
	}

	return pathMap[e]
}
