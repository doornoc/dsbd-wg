package peer

import (
	"fmt"
	"github.com/doornoc/dsbd-wg/pkg/core"
	"github.com/doornoc/dsbd-wg/pkg/core/db"
	"strings"
	"time"
)

func WgInit() error {
	db, err := db.ConnectDB()
	if err != nil {
		return err
	}

	dbSQL, err := db.DB()
	if err != nil {
		return fmt.Errorf("(%s)error: %s [%s]\n", time.Now(), "Failed to connect database.", err.Error())
	}
	defer dbSQL.Close()

	var peers []core.Client
	err = db.Find(&peers).Error
	if err != nil {
		return err
	}

	var wg_peers []Client

	for _, peer := range peers {
		wg_peers = append(wg_peers, Client{
			PublicKey:  peer.PublicKey,
			AllowedIps: strings.Split(peer.AllowedIps, ","),
			//Endpoint:   peer.Endpoint,
		})
	}

	_, err = WgAdd(wg_peers)
	if err != nil {
		return err
	}

	return nil
}
