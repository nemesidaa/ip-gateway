package main

import (
	"fmt"
	"testing"
)

func TestCalculation(t *testing.T) {
	cidr := 24 // Например, для /24
	subnetMask, err := CalculateSubnetMask(cidr)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	t.Logf("Маска подсети для /%d: %s\n", cidr, subnetMask)

	// Пример использования функции isIPInSubnet
	ip := "192.168.1.10"
	network := "192.168.1.0"
	ok, err := IsIPInSubnet(ip, network, cidr)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	if ok {
		t.Log("IP-адрес входит в подсеть")
	} else {
		t.Log("IP-адрес НЕ входит в подсеть")
	}
}
