package cielo_api30

type Environment struct {
	api, apiQuery string
}


func (env Environment) getApiUrl() string {
	return env.api
}

func (env Environment) getApiQueryURL() string {
	return env.apiQuery
}


var (
	SandboxEnvironment = Environment{api:"https://apisandbox.cieloecommerce.cielo_api30.com.br/",apiQuery:"https://apiquerysandbox.cieloecommerce.cielo_api30.com.br/"}
	ProductionEnvironment = Environment{api:"https://api.cieloecommerce.cielo_api30.com.br/",apiQuery:"https://apiquery.cieloecommerce.cielo_api30.com.br/"}
)