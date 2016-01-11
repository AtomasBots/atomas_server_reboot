package main
import (
	"os/exec"
	"time"
	"io/ioutil"
)

func main() {
	time.Sleep(time.Second * 300)
	for {
		if (isWifiOk()) {
			time.Sleep(time.Second)
		} else {
			time.Sleep(time.Second * 10)
			if (!isWifiOk()) {
				ioutil.WriteFile(time.Now().String(), []byte(time.Now().String()), 0644)
				exec.Command("sudo", "reboot").Start()
			}
		}
	}
}

func isWifiOk() bool {
	command := exec.Command("ping", "-c", "1", "192.168.0.1")
	command.Start()
	err := command.Wait()
	return err == nil
}