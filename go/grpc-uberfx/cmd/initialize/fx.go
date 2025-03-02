package initialize

import "go.uber.org/fx"

// Fx Modules untuk inject ke repository
var DBModule = fx.Provide(
	fx.Annotate(NewPostgresPrimaryDB, fx.ResultTags(`name:"postgrePrimaryDB"`)),
	fx.Annotate(NewPostgresStandbyDB, fx.ResultTags(`name:"postgreStandbyDB"`)),
	NewScyllaDB,
)
