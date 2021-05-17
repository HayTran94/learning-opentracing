package main

import (
	"fmt"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
	"io"
	"os"
	"time"
)

func initJaeger(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot int Jaeger: %v\n", err))
	}
	return tracer, closer
}

func main() {
	// get parameter
	if len(os.Args) != 2 {
		panic("ERROR: Expecting one argument")
	}
	helloTo := os.Args[1]
	helloStr := fmt.Sprintf("Hello, %s!", helloTo)

	// init tracing
	tracer, closer := initJaeger("hello-world")
	defer closer.Close()

	// start span
	span := tracer.StartSpan("say-hello")
	// set tag
	span.SetTag("hello-to", helloTo)
	// set log
	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", helloStr),
	)
	time.Sleep(time.Millisecond * 332)
	fmt.Println(helloStr)
	span.LogKV("event", "now")
	span.Finish()

}
