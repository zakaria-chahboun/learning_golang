package main

import (
	"fmt"

	"github.com/google/gousb"
)

func main() {
	gousbExample()
}

func gousbExample() {
	ctx := gousb.NewContext()
	defer ctx.Close()

	// List of USB devices
	devices, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		/*
			For example: if we just show the usb class, we will get this:
				hub
				wireless
				per-interface

			Our pos printer `desc.Class` is "per-interface"! Why?
			"per-interface" means that different interfaces of this device have different classes.
			So, we have to check all the classes of each interface! to get finally the "printer" class.
		*/
		for _, config := range desc.Configs {
			for _, inter := range config.Interfaces {
				for _, setting := range inter.AltSettings {
					// Check if the interface class of USB is `Printer`
					if setting.Class == gousb.ClassPrinter {
						return true // means append to devices
					}
				}
			}
		}
		return false // means not wanted
	})

	// Check errors
	if err != nil {
		panic(err)
	}

	// Check if printers exists
	if len(devices) == 0 {
		fmt.Println("No Printer Found!")
		return
	}

	// ---- finally list all printers info ----
	fmt.Println("Printers: ")
	for _, device := range devices {
		name, _ := device.Product()
		vid := device.Desc.Vendor
		pid := device.Desc.Product
		port := device.Desc.Port
		path := device.Desc.Path
		speed := device.Desc.Speed.String()
		configs := device.Desc.Configs
		serial, _ := device.SerialNumber()
		version := device.Desc.Device
		address := device.Desc.Address
		protocol := device.Desc.Protocol
		spec := device.Desc.Spec

		fmt.Println("name:", name)
		fmt.Println("vid:", vid)
		fmt.Println("pid:", pid)
		fmt.Println("port:", port)
		fmt.Println("path:", path)
		fmt.Println("speed:", speed)
		fmt.Println("configs:", configs)
		fmt.Println("serial:", serial)
		fmt.Println("version:", version)
		fmt.Println("address", address)
		fmt.Println("protocol", protocol)
		fmt.Println("spec", spec)
	}



}
