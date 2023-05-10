package config

type configKey string

const (
	AppInfoName                = "app_info_name"
	AppInfoStartupDuration     = "app_info_startupDuration"
	AppInfoVersion             = "app_info_version"
	DataSourcesPostgresHost    = "data_sources_postgres_host"
	DataSourcesPostgresName    = "data_sources_postgres_name"
	DataSourcesPostgresPort    = "data_sources_postgres_port"
	DataSourcesPostgresPwd     = "data_sources_postgres_pwd"
	DataSourcesPostgresUser    = "data_sources_postgres_user"
	DataSourcesPostgresSSLMode = "data_sources_postgres_sslmode"
	ServerRestAPIPort          = "server_rest_api_port"
	ServerTgApiKey             = "server_tg_api_key"
	WebhookHostURL             = "webhook_host_url"
)
