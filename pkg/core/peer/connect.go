package peer

import (
	"fmt"
	"github.com/doornoc/dsbd-wg/pkg/core/tool"
	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"net"
	"net/http"
)

func WgAdd(input []Client) (int, error) {
	var peers []wgtypes.PeerConfig

	client, err := wgctrl.New()
	if err != nil {
		tool.OutputLog(fmt.Sprintf("_[WgAdd] Code: %d, Error: %s", 0, err.Error()))
		return http.StatusInternalServerError, err
	}
	defer client.Close()

	_, err = client.Device("wg0")
	if err != nil {
		tool.OutputLog(fmt.Sprintf("_[WgAdd] Code: %d, Error: %s", 1, err.Error()))
		return http.StatusInternalServerError, err
	}

	for _, peer := range input {
		publicKey, err := wgtypes.ParseKey(peer.PublicKey)
		if err != nil {
			tool.OutputLog(fmt.Sprintf("_[WgAdd] Code: %d, Error: %s", 2, err.Error()))
			return http.StatusBadRequest, err
		}

		//addr, err := net.ResolveUDPAddr("udp", peer.Endpoint)
		//if err != nil {
		//	return http.StatusInternalServerError, err
		//}

		var ips []net.IPNet
		for _, allowedIp := range peer.AllowedIps {
			_, tmpIp, err := net.ParseCIDR(allowedIp)
			if err != nil {
				tool.OutputLog(fmt.Sprintf("_[WgAdd] Code: %d, Error: %s", 3, err.Error()))
				return http.StatusBadRequest, err
			}
			ips = append(ips, *tmpIp)
		}

		peers = append(peers, wgtypes.PeerConfig{
			PublicKey:  publicKey,
			AllowedIPs: ips,
			//Endpoint:   addr,
		})
	}

	deviceConfig := wgtypes.Config{Peers: peers}

	err = client.ConfigureDevice("wg0", deviceConfig)
	if err != nil {
		tool.OutputLog(fmt.Sprintf("_[WgAdd] Code: %d, Error: %s", 4, err.Error()))
		return http.StatusInternalServerError, err
	}

	return 0, nil
}

func WgDelete(inputPublicKey string) (int, error) {
	client, err := wgctrl.New()
	if err != nil {
		tool.OutputLog(fmt.Sprintf("_[WgDelete] Code: %d, Error: %s", 0, err.Error()))
		return http.StatusInternalServerError, err
	}
	defer client.Close()

	_, err = client.Device("wg0")
	if err != nil {
		tool.OutputLog(fmt.Sprintf("_[WgDelete] Code: %d, Error: %s", 1, err.Error()))
		return http.StatusInternalServerError, err
	}

	publicKey, err := wgtypes.ParseKey(inputPublicKey)
	if err != nil {
		tool.OutputLog(fmt.Sprintf("_[WgDelete] Code: %d, Error: %s", 2, err.Error()))
		return http.StatusBadRequest, err
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
		tool.OutputLog(fmt.Sprintf("_[WgDelete] Code: %d, Error: %s", 3, err.Error()))
		return http.StatusInternalServerError, err
	}

	return 0, nil
}

func WgAllDelete() (int, error) {
	client, err := wgctrl.New()
	if err != nil {
		tool.OutputLog(fmt.Sprintf("_[WgAllDelete] Code: %d, Error: %s", 0, err.Error()))
		return http.StatusInternalServerError, err
	}
	defer client.Close()

	device, err := client.Device("wg0")
	if err != nil {
		tool.OutputLog(fmt.Sprintf("_[WgAllDelete] Code: %d, Error: %s", 1, err.Error()))
		return http.StatusInternalServerError, err
	}

	var peers []wgtypes.PeerConfig

	for _, peer := range device.Peers {
		peers = append(peers, wgtypes.PeerConfig{
			PublicKey: peer.PublicKey,
			Remove:    true,
		})
	}

	err = client.ConfigureDevice("wg0", wgtypes.Config{Peers: peers})
	if err != nil {
		tool.OutputLog(fmt.Sprintf("_[WgAllDelete] Code: %d, Error: %s", 2, err.Error()))
		return http.StatusInternalServerError, err
	}

	return 0, nil
}

func WgPublicCheck(publicKey string) (int, error) {
	client, err := wgctrl.New()
	if err != nil {
		tool.OutputLog(fmt.Sprintf("_[WgPublicCheck] Code: %d, Error: %s", 0, err.Error()))
		return http.StatusInternalServerError, err
	}
	defer client.Close()

	device, err := client.Device("wg0")
	if err != nil {
		tool.OutputLog(fmt.Sprintf("_[WgPublicCheck] Code: %d, Error: %s", 1, err.Error()))
		return http.StatusInternalServerError, err
	}

	for _, peer := range device.Peers {
		if peer.PublicKey.String() == publicKey {
			tool.OutputLog(fmt.Sprintf("_[WgPublicCheck] Code: %d, Error: %s", 2, err.Error()))
			return http.StatusBadRequest, fmt.Errorf("%s", "public key is exist.")
		}
	}

	return 0, nil
}
