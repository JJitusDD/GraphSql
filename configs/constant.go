package configs

type Config struct {
	Port           string `json:"port" mapstructure:"port"`
	ProjectRootDir string `json:"project_root_dir" mapstructure:"project_root_dir"`
	ENV            string `json:"env" mapstructure:"env"`
}

const (
	GeneralErr = "Không xác định"

	//Name VAAccType
	OneTime = "Chuyển tiền một lần"
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
