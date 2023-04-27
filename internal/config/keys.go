package config

type configKey string

const (
	AppInfoName = "app_info_name"
	AppInfoStartupDuration = "app_info_startupDuration"
	AppInfoVersion = "app_info_version"
	DataSourcesPostgresDBHost = "data_sources_postgres_db_host"
	DataSourcesPostgresDBName = "data_sources_postgres_db_name"
	DataSourcesPostgresDBPort = "data_sources_postgres_db_port"
	DataSourcesPostgresDBPwd = "data_sources_postgres_db_pwd"
	DataSourcesPostgresDBUser = "data_sources_postgres_db_user"
	ServerRestAPIPort = "server_rest_api_port"
	ServerTgAPIKey = "server_tg_api_key"
	WebhookHostURL = "webhook_host_url"
)
