package flags

var AppName = "COBRA Core Service"
var AppVersion = "N/A"
var AppCommitHash = "N/A"

const (
	// Prefix for environment variables
	EnvPrefix = "COBRA"

	// Content Type Headers
	HeaderKeyContentType        = "Content-Type"
	HeaderKeyCOBRAAuthorization = "Authorization"
	HeaderKeyCOBRAAccessToken   = "X-COBRA-Access-Token"
	HeaderKeyCOBRATokenExpired  = "X-COBRA-Token-Expired"
	HeaderKeyCOBRASubject       = "X-COBRA-Subject"

	// Content Type Value
	ContentTypeJSON = "application/json; charset=utf-8"
	ContentTypeXML  = "application/xml; charset=utf-8"
	ContentTypeHTML = "text/html; charset=utf-8"

	// ACL
	ACLAuthenticatedUser      = ""
	ACLAuthenticatedAnonymous = "2"
	ACLEveryone               = "3"

	//status
	UserStatusActive = 1
	UserStatusBanned = 0
)
