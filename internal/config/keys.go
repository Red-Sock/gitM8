package config

type configKey string

const (
	AppInfoName             = "app_info_name"
	AppInfoStartupDuration  = "app_info_startupDuration"
	AppInfoVersion          = "app_info_version"
	DataSourcesPostgresHost = "data_sources_postgres_host"
	DataSourcesPostgresName = "data_sources_postgres_name"
	DataSourcesPostgresPort = "data_sources_postgres_port"
	DataSourcesPostgresPwd  = "data_sources_postgres_pwd"
	DataSourcesPostgresUser = "data_sources_postgres_user"
	ServerRestAPIPort       = "server_rest_api_port"
)
