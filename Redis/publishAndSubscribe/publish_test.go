package publishAndSubscribe

import "testing"

func TestPublish(t *testing.T) {
	Publish("test")
}

func TestSubscribe(t *testing.T) {
	Subscribe("test")
}

func TestOfficalCase(t *testing.T) {
	OfficalCase()
}

