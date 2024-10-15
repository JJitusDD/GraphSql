package configs

type Config struct {
	Port           string         `json:"port" mapstructure:"port"`
	ProjectRootDir string         `json:"project_root_dir" mapstructure:"project_root_dir"`
	ENV            string         `json:"env" mapstructure:"env"`
	Postgress      PostgresConfig `json:"postgres" mapstructure:"postgres"`
	Redis          RedisConfig    `json:"redis" mapstructure:"redis"`
	MongoURI       string         `json:"mongo_uri" mapstructure:"mongo_uri" gorm:"mongo_uri"`
}

const (
	GeneralErr = "Không xác định"
)

type Response struct {
	Meta Meta        `json:"meta" mapstructure:"meta"`
	Data interface{} `json:"data,omitempty" mapstructure:"data"`
}

type Meta struct {
	Code    int         `json:"code" mapstructure:"code"`
	Msg     string      `json:"msg" mapstructure:"msg"`
	Message interface{} `json:"message,omitempty" mapstructure:"message"`
}

type PostgresConfig struct {
	Host      string `json:"host" mapstructure:"host"`
	Port      int    `json:"port" mapstructure:"port"`
	DbName    string `json:"dbname" mapstructure:"dbname"`
	User      string `json:"user" mapstructure:"user"`
	Pass      string `json:"password" mapstructure:"password"`
	SSLMode   string `json:"sslmode" mapstructure:"sslmode"`
	Prefix    string `json:"prefix" mapstructure:"prefix"`
	DebugMode bool   `json:"debug_mode" mapstructure:"debug_mode"`
}

type RedisConfig struct {
	Addr     string `json:"addr" mapstructure:"addr"`
	Password string `json:"password" mapstructure:"password"`
	DB       int    `json:"db" mapstructure:"db"`
}
