package lifecycle

import (
	"fmt"
	"net/url"

	"github.com/lxc/lxd/shared/api"
)

// CertificateAction represents a lifecycle event action for Certificates.
type CertificateAction string

// All supported lifecycle events for Certificates.
const (
	CertificateCreated = CertificateAction(api.EventLifecycleCertificateCreated)
	CertificateDeleted = CertificateAction(api.EventLifecycleCertificateDeleted)
	CertificateUpdated = CertificateAction(api.EventLifecycleCertificateUpdated)
)

// Event creates the lifecycle event for an action on a Certificate.
func (a CertificateAction) Event(fingerprint string, requestor *api.EventLifecycleRequestor, ctx map[string]any) api.EventLifecycle {
	u := fmt.Sprintf("/1.0/certificates/%s", url.PathEscape(fingerprint))

	return api.EventLifecycle{
		Action:    string(a),
		Source:    u,
		Context:   ctx,
		Requestor: requestor,
	}
}
