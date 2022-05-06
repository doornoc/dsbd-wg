package peer

import (
	"github.com/gin-gonic/gin"
	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"net"
	"net/http"
)

func Add(c *gin.Context) {
	var input inputAdd
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	client, err := wgctrl.New()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	defer client.Close()

	_, err = client.Device("wg0")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	publicKey, err := wgtypes.ParseKey(input.PublicKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	addr, err := net.ResolveUDPAddr("udp", input.Endpoint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	var ips []net.IPNet
	for _, allowedIp := range input.AllowedIps {
		_, tmpIp, err := net.ParseCIDR(allowedIp)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		ips = append(ips, *tmpIp)
	}

	deviceConfig := wgtypes.Config{
		Peers: []wgtypes.PeerConfig{
			{
				PublicKey:  publicKey,
				AllowedIPs: ips,
				Endpoint:   addr,
			},
		},
	}

	err = client.ConfigureDevice("wg0", deviceConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func Delete(c *gin.Context) {
	var input inputDelete
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	client, err := wgctrl.New()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	defer client.Close()

	_, err = client.Device("wg0")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	publicKey, err := wgtypes.ParseKey(input.PublicKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	deviceConfig := wgtypes.Config{
		Peers: []wgtypes.PeerConfig{
			{
				PublicKey: publicKey,
				Remove:    true,
			},
		},
	}

	err = client.ConfigureDevice("wg0", deviceConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func Get(c *gin.Context) {
	client, err := wgctrl.New()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	defer client.Close()

	device, err := client.Device("wg0")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	var d []data
	for _, peer := range device.Peers {
		var ips []string
		for _, ip := range peer.AllowedIPs {
			ips = append(ips, ip.String())
		}
		d = append(d, data{
			AllowedIps:        ips,
			LastHandshakeTime: peer.LastHandshakeTime,
			Endpoint:          peer.Endpoint.String(),
			PresharedKey:      peer.PresharedKey.String(),
			PublicKey:         peer.PublicKey.String(),
			ReceiveBytes:      peer.ReceiveBytes,
			TransmitBytes:     peer.TransmitBytes,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    d,
	})
}
