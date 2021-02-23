package dashboard

import (
	"math"

	traefikdynamic "github.com/containous/traefik/v2/pkg/config/dynamic"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/http"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/http/config/dynamic"
)

func AddDynamicConfig(cfg *dynamic.Configuration, middlewares []string) {
	cfg.HTTP.Routers["dashboard"] = &dynamic.Router{
		Router: &traefikdynamic.Router{
			EntryPoints: []string{http.DefaultHTTPAppEntryPoint},
			Service:     "dashboard",
			Priority:    math.MaxInt32,
			Rule:        "PathPrefix(`/api`) || PathPrefix(`/dashboard`)",
			Middlewares: append(middlewares, "strip-api"),
		},
	}

	cfg.HTTP.Middlewares["strip-api"] = &dynamic.Middleware{
		Middleware: &traefikdynamic.Middleware{
			StripPrefixRegex: &traefikdynamic.StripPrefixRegex{
				Regex: []string{"/api"},
			},
		},
	}

	// Dashboard API
	cfg.HTTP.Services["dashboard"] = &dynamic.Service{
		Dashboard: &dynamic.Dashboard{},
	}
}