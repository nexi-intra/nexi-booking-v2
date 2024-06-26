/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: true
---
*/
package magicapp

import (
	"github.com/nats-io/nats.go/micro"
	"github.com/nexi-intra/nexi-booking/services"
)

func RegisterServiceCoreEndpoints(root micro.Group) {
	root.AddEndpoint("app", micro.HandlerFunc(services.HandleAppRequests))
}
