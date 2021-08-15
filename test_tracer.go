package testtracer

import (
	"log"
	"sync"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func init() {
	var once sync.Once
	once.Do(func() {
		cfg := setupJaeger()
		tracer, _, err := cfg.NewTracer()
		if err != nil {
			log.Printf("Could not initialize jaeger tracer: %s", err.Error())
			return
		}
		opentracing.SetGlobalTracer(tracer)
	})
}

func setupJaeger() *jaegercfg.Configuration {
	// Sample configuration for testing. Use constant sampling to sample every trace
	// and enable LogSpan to log every span via configured Logger.
	jaegercfg.FromEnv()
	cfg := jaegercfg.Configuration{
		ServiceName: "testtracer",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}
	return &cfg
}
