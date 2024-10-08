package config

type Logger struct {
	LogLevel     string `mapstructure:"level"` // 日志级别 debugger输出全部SQL dev 输出错误日志 release
	Prefix       string `mapstructure:"prefix"`
	ShowLine     bool   `mapstructure:"show_line"`      // 是否显示行号
	LogTOConsole bool   `mapstructure:"log_to_console"` // 是否输出到控制台
	LogToFile    bool   `mapstructure:"log_to_file"`
	FilePath     string `mapstructure:"file_path"`
}
