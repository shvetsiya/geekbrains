package config

import (
	"fmt"
	"net/url"
	"strings"
)

type Validator struct {
	err error
}

func (v *Validator) IsPortValid(value, lbound, rbound uint16) bool {
	if v.err != nil {
		return false
	}
	if value <= lbound || value >= rbound {
		v.err = fmt.Errorf("The port value must be in the range: %d-%d", lbound, rbound)
		return false
	}
	return true
}

func (v *Validator) IsDbURLValid(str, pre string) bool {
	if v.err != nil {
		return false
	}

	if !strings.HasPrefix(str, pre) {
		v.err = fmt.Errorf("Db route has incorrect URL")
		return false
	}
	return true
}

func (v *Validator) IsURLValid(urlPath, host string) bool {
	if v.err != nil {
		return false
	}

	u, err := url.Parse(urlPath)
	if err != nil {
		v.err = fmt.Errorf("Can not parse URL %v", urlPath)
	}

	if !strings.HasPrefix(u.Host, host) {
		v.err = fmt.Errorf("URL %v has different host %v", urlPath, host)
	}
	return true
}

func (v *Validator) IsKafkaBrokerValid(str, pre string) bool {
	if v.err != nil {
		return false
	}

	if !strings.HasPrefix(str, pre) {
		v.err = fmt.Errorf("Kafka broker is incorrect %v", str)
		return false
	}
	return true
}

func (v *Validator) IsValid() bool {
	return v.err == nil
}
