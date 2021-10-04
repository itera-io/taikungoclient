
// add support for "application/problem+json" media type
transport.Consumers["application/problem+json"] = runtime.JSONConsumer()
transport.Producers["application/problem+json"] = runtime.JSONProducer()

// add support for "application/*+json" media type
transport.Consumers["application/*+json"] = runtime.JSONConsumer()
transport.Producers["application/*+json"] = runtime.JSONProducer()

// add support for "application/json-patch+json" media type
transport.Consumers["application/json-patch+json"] = runtime.JSONConsumer()
transport.Producers["application/json-patch+json"] = runtime.JSONProducer()

// add support for "text/json" media type
transport.Consumers["text/json"] = runtime.JSONConsumer()
transport.Producers["text/json"] = runtime.JSONProducer()

