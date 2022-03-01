package romulus

import "context"

type Lifecycle interface {
	Name() string
	Start(context.Context) error
	Stop(context.Context) error
}
