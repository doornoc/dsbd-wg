package peer

import (
	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"net"
	"testing"
)

func TestGetWireGuard(t *testing.T) {
	t.Log("=======TEST=======")
	client, err := wgctrl.New()
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer client.Close()

	device, err := client.Device("wg0")
	if err != nil {
		t.Fatalf(err.Error())
	}

	for _, tmp := range device.Peers {
		t.Log("==========================")
		t.Log("AllowedIPs: ", tmp.AllowedIPs)
		t.Log("LastHandshakeTime: ", tmp.LastHandshakeTime)
		t.Log("Endpoint: ", tmp.Endpoint)
		t.Log("PresharedKey: ", tmp.PresharedKey)
		t.Log("PublicKey: ", tmp.PublicKey.String())
		t.Log("ReceiveBytes: ", tmp.ReceiveBytes)
		t.Log("TransmitBytes: ", tmp.TransmitBytes)
	}
}

func TestAddWireGuard(t *testing.T) {
	t.Log("=======ADD TEST=======")
	client, err := wgctrl.New()
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer client.Close()

	_, err = client.Device("wg0")
	if err != nil {
		t.Fatalf(err.Error())
	}

	key, _ := wgtypes.ParseKey("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")

	addr, err := net.ResolveUDPAddr("udp", "x.x.x.x:xxxx")
	if err != nil {
		t.Fatalf(err.Error())
	}

	_, v4, err := net.ParseCIDR("x.x.x.x/32")
	if err != nil {
		t.Fatalf(err.Error())
	}

	_, v6, err := net.ParseCIDR("xxxx:xxxx:xxxx:xxxx::x/128")
	if err != nil {
		t.Fatalf(err.Error())
	}

	deviceConfig := wgtypes.Config{
		Peers: []wgtypes.PeerConfig{
			{
				PublicKey:  key,
				AllowedIPs: []net.IPNet{*v4, *v6},
				Endpoint:   addr,
			},
		},
	}

	err = client.ConfigureDevice("wg0", deviceConfig)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestRemoveWireGuard(t *testing.T) {
	t.Log("=======REMOVE TEST=======")
	client, err := wgctrl.New()
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer client.Close()

	_, err = client.Device("wg0")
	if err != nil {
		t.Fatalf(err.Error())
	}

	key, _ := wgtypes.ParseKey("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	t.Log("key: ", key)

	deviceConfig := wgtypes.Config{
		Peers: []wgtypes.PeerConfig{
			{
				PublicKey: key,
				Remove:    true,
			},
		},
	}

	err = client.ConfigureDevice("wg0", deviceConfig)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

//func TestGenerateKey(t *testing.T) {
//	priv, err := wgtypes.GeneratePrivateKey()
//	if err != nil {
//		t.Fatalf("failed to generate private key: %v", err)
//	}
//	t.Log(priv.String())
//}
