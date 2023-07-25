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
const (
	InsertTodoListBindingKey = "todolist.action.insert"
	EditTodoListBindingKey   = "todolist.action.edit"
	RemoveTodoListBindingKey = "todolist.action.remove"

	InsertTodoListQueueName = "todolist.action.insert"
	EditTodoListQueueName   = "todolist.action.edit"
	RemoveTodoListQueueName = "todolist.action.remove"

	InsertTodoItemBindingKey = "todoitem.action.insert"
	EditTodoItemBindingKey   = "todoitem.action.edit"
	RemoveTodoItemBindingKey = "todoitem.action.remove"

	InsertTodoItemQueueName = "todoitem.action.insert"
	EditTodoItemQueueName   = "todoitem.action.edit"
	RemoveTodoItemQueueName = "todoitem.action.remove"
)
