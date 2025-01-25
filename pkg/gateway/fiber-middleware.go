package gateway

import (
	"github.com/gofiber/fiber/v2"
	calculations "github.com/nemesidaa/ip-gateway/pkg/calc"
)

// ! Config input formats: ip/cidr only
// Firstly you should run UploadConfig
func V4Gateway(c *fiber.Ctx) error {
	ip := c.Get(conf.IpHeaderName)
	if ip == "" {
		return ErrNoIp
	}
	pass := false
	for _, net := range conf.Networks {
		ok, err := calculations.IsIPInSubnet(ip, net.Subnet, net.Mask)
		if err != nil {
			return ErrFailedToCalculateIP
		}
		if ok {
			pass = true
			break
		}
	}
	if !pass {
		return ErrFailedToPass
	}

	return c.Next()
}
