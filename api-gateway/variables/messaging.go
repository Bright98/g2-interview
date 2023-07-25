package variables

// publisher
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

const (
	//user actions
	InsertUserBindingKey = "user.action.insert"
	EditUserBindingKey   = "user.action.edit"
	RemoveUserBindingKey = "user.action.remove"

	//user events
	InsertUserQueueName = "user.event.insert"
	EditUserQueueName   = "user.event.edit"
	RemoveUserQueueName = "user.event.remove"
)
