package setting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}

type AppSettingS struct {
	DefaultPageSize         int
	MaxPageSize             int
	NormalUserDeviceUplimit int
	VipUserDeviceUplimit    int
	LogSavePath             string
	LogFileName             string
	LogFileExt              string
}

type DatabaseSettingS struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type JWTSettingS struct {
	AppSecret string
	Issuer    string
	Expire    time.Duration
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
