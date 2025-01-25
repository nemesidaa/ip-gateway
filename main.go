package main

import (
	"fmt"

	"github.com/nemesidaa/ip-gateway/pkg/gateway"
)

func main() {
	// Пример использования функции calculateSubnetMask
	cidr := 24 // Например, для /24
	subnetMask, err := gateway.CalculateSubnetMask(cidr)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Printf("Маска подсети для /%d: %s\n", cidr, subnetMask)

	// Пример использования функции isIPInSubnet
	ip := "192.168.1.10"
	network := "192.168.1.0"
	ok, err := gateway.IsIPInSubnet(ip, network, cidr)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	if ok {
		fmt.Printf("IP-адрес %s входит в подсеть %s/%d\n", ip, network, cidr)
	} else {
		fmt.Printf("IP-адрес %s НЕ входит в подсеть %s/%d\n", ip, network, cidr)
	}
}
