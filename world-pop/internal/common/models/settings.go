package models

type Settings struct {
	DataFilePath     	string 		`yaml:"data_file_path"`
	MinimumLogLevel  	string 		`yaml:"minimum_log_level"`
	UpdaterServerUrl 	string 		`yaml:"updater_server_url"`
	EnableLogging    	bool   		`yaml:"enable_logging"`
	AutoUpdate			bool		`yaml:"auto_update"`
	SettingFilePath		string		`yaml:"-"`
}