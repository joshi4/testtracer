package testtracer

import (
	"io"
	"log"
	"sync"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

var closeFn io.Closer

func init() {
	var once sync.Once
	once.Do(func() {
		cfg := setupJaeger()
		tracer, closer, err := cfg.NewTracer()
		if err != nil {
			log.Printf("Could not initialize jaeger tracer: %s", err.Error())
			return
		}
		closeFn = closer
		opentracing.SetGlobalTracer(tracer)
	})
}

func Close() error {
	println("calling close")
	if closeFn != nil {
		return closeFn.Close()
	}
	return nil
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
