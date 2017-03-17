package mailgun

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"time"
)

type EventType uint8

const (
	EventUnknown EventType = iota
	EventAccepted
	EventRejected
	EventDelivered
	EventFailed
	EventOpened
	EventClicked
	EventUnsubscribed
	EventComplained
	EventStored
)

var eventTypes = []string{
	"unknown",
	"accepted",
	"rejected",
	"delivered",
	"failed",
	"opened",
	"clicked",
	"unsubscribed",
	"complained",
	"stored",
}

func (et EventType) String() string {
	return eventTypes[et]
}

// MarshalText satisfies TextMarshaler
func (et EventType) MarshalText() ([]byte, error) {
	return []byte(et.String()), nil
}

// UnmarshalText satisfies TextUnmarshaler
func (et *EventType) UnmarshalText(text []byte) error {
	enum := string(text)
	for i := 0; i < len(eventTypes); i++ {
		if enum == eventTypes[i] {
			*et = EventType(i)
			return nil
		}
	}
	return fmt.Errorf("unknown event type '%s'", enum)
}

type TimestampNano time.Time

// MarshalText satisfies TextMarshaler
func (tn TimestampNano) MarshalText() ([]byte, error) {
	t := time.Time(tn)
	v := float64(t.Unix()) + float64(t.Nanosecond())/float64(time.Nanosecond)
	return json.Marshal(v)
}

// UnmarshalText satisfies TextUnmarshaler
func (tn *TimestampNano) UnmarshalText(text []byte) error {
	v, err := strconv.ParseFloat(string(text), 64)
	if err == nil {
		*tn = TimestampNano(time.Unix(0, int64(v*float64(time.Second))))
	}
	return err
}

type IP net.IP

// MarshalText satisfies TextMarshaler
func (i IP) MarshalText() ([]byte, error) {
	return []byte(net.IP(i).String()), nil
}

// UnmarshalText satisfies TextUnmarshaler
func (i *IP) UnmarshalText(text []byte) error {
	v := net.ParseIP(string(text))
	if v != nil {
		*i = IP(v)
	}
	return nil
}

type Method uint8

const (
	MethodUnknown Method = iota
	MethodSMTP
	MethodAPI
)

var methods = []string{
	"unknown",
	"smtp",
	"api",
}

func (m Method) String() string {
	return methods[m]
}

// MarshalText satisfies TextMarshaler
func (m Method) MarshalText() ([]byte, error) {
	return []byte(m.String()), nil
}

// UnmarshalText satisfies TextUnmarshaler
func (m *Method) UnmarshalText(text []byte) error {
	enum := string(text)
	for i := 0; i < len(methods); i++ {
		if enum == methods[i] {
			*m = Method(i)
			return nil
		}
	}
	return fmt.Errorf("unknown event method '%s'", enum)
}

type EventSeverity uint8

const (
	SeverityUnknown EventSeverity = iota
	SeverityTemporary
	SeverityPermanent
	SeverityInternal
)

var severities = []string{
	"unknown",
	"permanent",
	"temporary",
	"internal",
}

func (es EventSeverity) String() string {
	return severities[es]
}

// MarshalText satisfies TextMarshaler
func (es EventSeverity) MarshalText() ([]byte, error) {
	return []byte(es.String()), nil
}

// UnmarshalText satisfies TextUnmarshaler
func (es *EventSeverity) UnmarshalText(text []byte) error {
	enum := string(text)
	for i := 0; i < len(severities); i++ {
		if enum == severities[i] {
			*es = EventSeverity(i)
			return nil
		}
	}
	return fmt.Errorf("unknown event severity '%s'", enum)
}

type EventReason uint8

const (
	ReasonUnknown EventReason = iota
	ReasonGeneric
	ReasonBounce
	ReasonESPBlock
	ReasonSuppressBounce
	ReasonSuppressComplaint
	ReasonSuppressUnsubscribe
	ReasonOld
)

var eventReasons = []string{
	"unknown",
	"generic",
	"bounce",
	"espblock",
	"suppress-bounce",
	"suppress-complaint",
	"suppress-unsubscribe",
	"old",
}

func (er EventReason) String() string {
	return eventReasons[er]
}

// MarshalText satisfies TextMarshaler
func (er EventReason) MarshalText() ([]byte, error) {
	return []byte(er.String()), nil
}

// UnmarshalText satisfies TextUnmarshaler
func (er *EventReason) UnmarshalText(text []byte) error {
	enum := string(text)
	for i := 0; i < len(eventReasons); i++ {
		if enum == eventReasons[i] {
			*er = EventReason(i)
			return nil
		}
	}
	return fmt.Errorf("unknown event reason '%s'", enum)
}

type ClientType uint

const (
	ClientUnknown ClientType = iota
	ClientMobileBrowser
	ClientBrowser
	ClientEmail
)

var clientTypes = []string{
	"unknown",
	"mobile browser",
	"browser",
	"email client",
}

func (ct ClientType) String() string {
	return clientTypes[ct]
}

// MarshalText satisfies TextMarshaler
func (ct ClientType) MarshalText() ([]byte, error) {
	return []byte(ct.String()), nil
}

// UnmarshalText satisfies TextUnmarshaler
func (ct *ClientType) UnmarshalText(text []byte) error {
	enum := string(text)
	for i := 0; i < len(clientTypes); i++ {
		if enum == clientTypes[i] {
			*ct = ClientType(i)
			return nil
		}
	}
	return fmt.Errorf("unknown client type '%s'", enum)
}

type DeviceType uint

const (
	DeviceUnknown DeviceType = iota
	DeviceMobileBrowser
	DeviceBrowser
	DeviceEmail
)

var deviceTypes = []string{
	"unknown",
	"desktop",
	"mobile",
	"tablet",
}

func (ct DeviceType) String() string {
	return deviceTypes[ct]
}

// MarshalText satisfies TextMarshaler
func (ct DeviceType) MarshalText() ([]byte, error) {
	return []byte(ct.String()), nil
}

// UnmarshalText satisfies TextUnmarshaler
func (ct *DeviceType) UnmarshalText(text []byte) error {
	enum := string(text)
	for i := 0; i < len(deviceTypes); i++ {
		if enum == deviceTypes[i] {
			*ct = DeviceType(i)
			return nil
		}
	}
	return fmt.Errorf("unknown device type '%s'", enum)
}

type TransportMethod uint

const (
	TransportUnknown TransportMethod = iota
	TransportHTTP
	TransportSMTP
)

var transportMethods = []string{
	"unknown",
	"http",
	"smtp",
}

func (tm TransportMethod) String() string {
	return transportMethods[tm]
}

// MarshalText satisfies TextMarshaler
func (tm TransportMethod) MarshalText() ([]byte, error) {
	return []byte(tm.String()), nil
}

// UnmarshalText satisfies TextUnmarshaler
func (tm *TransportMethod) UnmarshalText(text []byte) error {
	enum := string(text)
	for i := 0; i < len(transportMethods); i++ {
		if enum == transportMethods[i] {
			*tm = TransportMethod(i)
			return nil
		}
	}
	return fmt.Errorf("unknown transport method '%s'", enum)
}
