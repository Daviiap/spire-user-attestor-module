package domain

type UserInfo struct {
	UnixInfo   UnixInfo
	BasicAuth  BasicAuth
	External   External
	Biometrics Biometrics
}

type UnixInfo struct {
	UID           string
	User          string
	GID           string
	Group         string
	Supplementary []Supplementary
}

type Supplementary struct {
	GID   string
	Group string
}

type BasicAuth struct {
	User     string
	Password string
}

type External struct{}

type Biometrics struct{}
