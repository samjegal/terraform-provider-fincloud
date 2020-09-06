package vpc

// FINCLOUD_APACHE_NO_VERSION

// UserAgent returns the UserAgent string to use when sending http.Requests.
func UserAgent() string {
	return "Azure-SDK-For-Go/" + Version() + " vpc/1.0.0"
}

// Version returns the semantic version (see http://semver.org) of the client.
func Version() string {
	return "v0.0.1"
}
