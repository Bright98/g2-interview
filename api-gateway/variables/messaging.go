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

	//todo list actions
	InsertTodoListBindingKey = "todolist.action.insert"
	EditTodoListBindingKey   = "todolist.action.edit"
	RemoveTodoListBindingKey = "todolist.action.remove"

	//todo list events
	InsertTodoListQueueName = "todolist.event.insert"
	EditTodoListQueueName   = "todolist.event.edit"
	RemoveTodoListQueueName = "todolist.event.remove"

	//todo item actions
	InsertTodoItemBindingKey = "todoitem.action.insert"
	EditTodoItemBindingKey   = "todoitem.action.edit"
	RemoveTodoItemBindingKey = "todoitem.action.remove"

	//todo item events
	InsertTodoItemQueueName = "todoitem.event.insert"
	EditTodoItemQueueName   = "todoitem.event.edit"
	RemoveTodoItemQueueName = "todoitem.event.remove"
)
