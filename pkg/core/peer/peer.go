package peer

import (
	"fmt"
	"github.com/doornoc/dsbd-wg/pkg/core"
	"github.com/doornoc/dsbd-wg/pkg/core/db"
	"github.com/doornoc/dsbd-wg/pkg/core/tool"
	"github.com/gin-gonic/gin"
	"golang.zx2c4.com/wireguard/wgctrl"
	"net/http"
	"strings"
)

func Add(c *gin.Context) {
	var input Client
	err := c.BindJSON(&input)
	if err != nil {
		tool.OutputLog(fmt.Sprintf("[Add] Code: %d, Error: %s", 0, err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	errCode, err := WgPublicCheck(input.PublicKey)
	if err != nil {
		tool.OutputLog(fmt.Sprintf("[Add] Code: %d, Error: %s", 1, err.Error()))
		c.JSON(errCode, gin.H{"message": err.Error()})
		return
	}

	errCode, err = WgAdd([]Client{input})
	if err != nil {
		tool.OutputLog(fmt.Sprintf("[Add] Code: %d, Error: %s", 2, err.Error()))
		c.JSON(errCode, gin.H{"message": err.Error()})
		return
	}
	err = db.Add([]*core.Client{
		{
			PublicKey:  input.PublicKey,
			AllowedIps: strings.Join(input.AllowedIps, ","),
			//Endpoint:   input.Endpoint,
		},
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func Delete(c *gin.Context) {
	var input inputDelete
	err := c.BindJSON(&input)
	if err != nil {
		tool.OutputLog(fmt.Sprintf("[Delete] Code: %d, Error: %s", 0, err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	errCode, err := WgDelete(input.PublicKey)
	if err != nil {
		tool.OutputLog(fmt.Sprintf("[Delete] Code: %d, Error: %s", 1, err.Error()))
		c.JSON(errCode, gin.H{"message": err.Error()})
		return
	}

	err = db.Delete(input.PublicKey)
	if err != nil {
		tool.OutputLog(fmt.Sprintf("[Delete] Code: %d, Error: %s", 2, err.Error()))
		c.JSON(errCode, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func AllDelete(c *gin.Context) {
	errCode, err := WgAllDelete()
	if err != nil {
		tool.OutputLog(fmt.Sprintf("[AllDelete] Code: %d, Error: %s", 0, err.Error()))
		c.JSON(errCode, gin.H{"message": err.Error()})
		return
	}

	err = db.DeleteAll()
	if err != nil {
		tool.OutputLog(fmt.Sprintf("[AllDelete] Code: %d, Error: %s", 1, err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

// public_keyから削除
func Put(c *gin.Context) {
	var input Edit
	err := c.BindJSON(&input)
	if err != nil {
		tool.OutputLog(fmt.Sprintf("[Put] Code: %d, Error: %s", 0, err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	errCode, err := WgPublicCheck(input.Client.PublicKey)
	if err != nil {
		tool.OutputLog(fmt.Sprintf("[Put] Code: %d, Error: %s", 1, err.Error()))
		c.JSON(errCode, gin.H{"message": err.Error()})
		return
	}

	errCode, err = WgDelete(input.OldPublicKey)
	if err != nil {
		tool.OutputLog(fmt.Sprintf("[Put] Code: %d, Error: %s", 2, err.Error()))
		c.JSON(errCode, gin.H{"message": err.Error()})
		return
	}

	err = db.Delete(input.OldPublicKey)
	if err != nil {
		tool.OutputLog(fmt.Sprintf("[Put] Code: %d, Error: %s", 3, err.Error()))
		c.JSON(errCode, gin.H{"message": err.Error()})
		return
	}

	errCode, err = WgAdd([]Client{input.Client})
	if err != nil {
		tool.OutputLog(fmt.Sprintf("[Put] Code: %d, Error: %s", 4, err.Error()))
		c.JSON(errCode, gin.H{"message": err.Error()})
		return
	}

	err = db.Add([]*core.Client{
		{
			PublicKey:  input.Client.PublicKey,
			AllowedIps: strings.Join(input.Client.AllowedIps, ","),
			//Endpoint:   input.Client.Endpoint,
		},
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func Overwrite(c *gin.Context) {
	var input Clients
	err := c.BindJSON(&input)
	if err != nil {
		tool.OutputLog(fmt.Sprintf("[Overwrite] Code: %d, Error: %s", 0, err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	errCode, err := WgAllDelete()
	if err != nil {
		tool.OutputLog(fmt.Sprintf("[Overwrite] Code: %d, Error: %s", 1, err.Error()))
		c.JSON(errCode, gin.H{"message": err.Error()})
		return
	}

	err = db.DeleteAll()
	if err != nil {
		tool.OutputLog(fmt.Sprintf("[Overwrite] Code: %d, Error: %s", 2, err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	errCode, err = WgAdd(input.Clients)
	if err != nil {
		tool.OutputLog(fmt.Sprintf("[Overwrite] Code: %d, Error: %s", 3, err.Error()))
		c.JSON(errCode, gin.H{"message": err.Error()})
		return
	}

	var clients []*core.Client
	for _, peer := range input.Clients {
		clients = append(clients, &core.Client{
			PublicKey:  peer.PublicKey,
			AllowedIps: strings.Join(peer.AllowedIps, ","),
			//Endpoint:   peer.Endpoint,
		})
	}

	err = db.Add(clients)

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func Check(c *gin.Context) {
	var input Client
	err := c.BindJSON(&input)
	if err != nil {
		tool.OutputLog(fmt.Sprintf("[Check] Code: %d, Error: %s", 0, err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	errCode, err := WgPublicCheck(input.PublicKey)
	if err != nil {
		tool.OutputLog(fmt.Sprintf("[Check] Code: %d, Error: %s", 1, err.Error()))
		c.JSON(errCode, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func Get(c *gin.Context) {
	client, err := wgctrl.New()
	if err != nil {
		tool.OutputLog(fmt.Sprintf("[Get] Code: %d, Error: %s", 0, err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	defer client.Close()

	device, err := client.Device("wg0")
	if err != nil {
		tool.OutputLog(fmt.Sprintf("[Get] Code: %d, Error: %s", 1, err.Error()))
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
