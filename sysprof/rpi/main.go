package main

import (
	"fmt"
	"math"
	"os/exec"
	"strconv"
	"strings"
)

type Measurement interface {
	Voltage() int
	Temperature() int
	NetBand()
}

type Command struct {
	Cmd  string
	Args []string
}

func main() {
	temp_var := Command{
		Cmd:  "vcgencmd",
		Args: []string{"measure_temp"},
	}

	volt_var := Command{
		Cmd:  "vcgencmd",
		Args: []string{"get_throttled"},
	}
	temp_var.Temperature()
	volt_var.Voltage()
}

func (c *Command) Voltage() int {
	cmd := exec.Command(c.Cmd, c.Args...)
	op, err := cmd.CombinedOutput()
	if err != nil {
		return -1
	}
	replacer := strings.NewReplacer("throttled=", "", "\n", "")
	status_code := replacer.Replace(string(op))

	if status_code == "0x0" {
		return 0
	} else {
		return 1
	}
}

func (c *Command) Temperature() int {
	cmd := exec.Command(c.Cmd, c.Args...)
	op, err := cmd.CombinedOutput()
	if err != nil {
		return -1
	}
	replacer := strings.NewReplacer("temp=", "", "'C", "", "\n", "")
	str_temp := replacer.Replace(string(op))
	temp, err := strconv.ParseFloat(str_temp, 32)
	if err != nil {
		return -1
	}
	if math.Round(temp) <= 85 {
		return 0
	} else {
		return 1
	}
}

func (c *Command) NetBand() string {
	cmd := exec.Command(c.Cmd, c.Args...)
	op, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	return string(op)
}
