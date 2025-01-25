package gateway

import (
	"fmt"
	"net"
)

func CalculateSubnetMask(cidr int) (string, error) {
	if cidr < 0 || cidr > 32 {
		return "", fmt.Errorf("CIDR должен быть в диапазоне от 0 до 32")
	}

	mask := net.CIDRMask(cidr, 32)

	if mask == nil {
		return "", fmt.Errorf("не удалось создать маску подсети")
	}

	return net.IP(mask).String(), nil
}

// Функция для проверки, входит ли IP-адрес в подсеть
func IsIPInSubnet(ip string, network string, cidr int) (bool, error) {
	// Преобразуем IP-адрес и сеть в тип net.IP
	ipAddr := net.ParseIP(ip)
	if ipAddr == nil {
		return false, fmt.Errorf("неверный формат IP-адреса: %s", ip)
	}

	// Создаем сеть на основе CIDR
	_, subnet, err := net.ParseCIDR(fmt.Sprintf("%s/%d", network, cidr))
	if err != nil {
		return false, fmt.Errorf("неверный формат сети или CIDR: %v", err)
	}

	// Проверяем, входит ли IP-адрес в подсеть
	return subnet.Contains(ipAddr), nil
}
