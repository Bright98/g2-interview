package variables

// subscriber
const (
	ExchangeName       = "rabbit-exchange"
	ExchangeKind       = "topic"
	ExchangeDurable    = true
	ExchangeAutoDelete = false
	ExchangeInternal   = false
	ExchangeNoWait     = false

	QueueDurable    = true
	QueueAutoDelete = false
	QueueExclusive  = false
	QueueNoWait     = false

	PublishMandatory = false
	PublishImmediate = false

	PrefetchCount  = 1
	PrefetchSize   = 0
	PrefetchGlobal = false

	ConsumeAutoAck   = false
	ConsumeExclusive = false
	ConsumeNoLocal   = false
	ConsumeNoWait    = false
)

// queue / key names
const ()
