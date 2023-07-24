package variables

// subscriber
const (
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

// exchange name
const (
	ExchangeName = "rabbit-exchange"
)

// queue / key names
const (
	InsertUserQueueName = "user.insert"
	EditUserQueueName   = "user.dit"
	RemoveUserQueueName = "user.remove"
)
