
A simple LED Blinker in [GO](https://golang.org/) using the [EMBD library](https://github.com/kidoman/embd) to drive the GPIO pins.

This blinker program was testing using a setup with a Raspberry Pi Zero, a Red LED (about 1.7 voltage drop) driven by pin GPIO_4 and a current limiting resistor of 100 Ohms.

The project is built using [GB](https://getgb.io/). See [build4arm.sh](build4arm.sh).

**Dependencies**
- [github.com/kidoman/embd](https://github.com/kidoman/embd)
- [github.com/golang/glog](https://github.com/golang/glog)

_Disclaimer: This example was glued using the following references:_
- [\[1\] github.com/kidoman/embd/blob/master/samples/fullblinker.go](https://github.com/kidoman/embd/blob/master/samples/fullblinker.go);
- [\[2\] github.com/kidoman/embdgithub.com/kidoman/embd](https://github.com/kidoman/embd).
