package locust

import "github.com/amila-ku/locust-client"


func Swarm() {
	client, err := New(locusturl)
	if err != nil {
		return err
	}
	status, err := client.Start(5, 1)
	if err != nil {
		return err
	}
}