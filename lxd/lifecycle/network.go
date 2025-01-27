package lifecycle

import (
	"fmt"
	"net/url"

	"github.com/lxc/lxd/lxd/project"
	"github.com/lxc/lxd/shared/api"
)

// Internal copy of the network interface.
type network interface {
	Name() string
	Project() string
}

// NetworkAction represents a lifecycle event action for network devices.
type NetworkAction string

// All supported lifecycle events for network devices.
const (
	NetworkCreated = NetworkAction(api.EventLifecycleNetworkCreated)
	NetworkDeleted = NetworkAction(api.EventLifecycleNetworkDeleted)
	NetworkUpdated = NetworkAction(api.EventLifecycleNetworkUpdated)
	NetworkRenamed = NetworkAction(api.EventLifecycleNetworkRenamed)
)

// Event creates the lifecycle event for an action on a network device.
func (a NetworkAction) Event(n network, requestor *api.EventLifecycleRequestor, ctx map[string]any) api.EventLifecycle {
	u := fmt.Sprintf("/1.0/networks/%s", url.PathEscape(n.Name()))
	if n.Project() != project.Default {
		u = fmt.Sprintf("%s?project=%s", u, url.QueryEscape(n.Project()))
	}

	return api.EventLifecycle{
		Action:    string(a),
		Source:    u,
		Context:   ctx,
		Requestor: requestor,
	}
}
