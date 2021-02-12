package types

// SecretStore is used to store secrets
type SecretStore interface {
	LoadSecrets(scope string) error
	SaveSecrets(scope string) error
	SetSecret(name string, value interface{}, scope string) error
	SetSecrets(data map[string]interface{}, scope string) error
	GetSecret(name, scope string) (interface{}, error)
	GetSecrets(scope string) (map[string]interface{}, error)
	ListSecrets(scope string) ([]string, error)
	SetAppName(name string) error
	ListAppNames() ([]string, error)
	GetCurrentVersion(scope string) (int64, error)
	GetLatestVersion(scope string) (int64, error)
}