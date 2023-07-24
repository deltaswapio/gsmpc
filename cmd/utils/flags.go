package utils

import (
	//"crypto/ecdsa"
	//"fmt"
	//"io"
	//"io/ioutil"
	//"math/big"
	//"os"
	//"path/filepath"
	//"strconv"
	//"strings"
	//"text/tabwriter"
	//"text/template"
	//"time"

	//"github.com/deltaswapio/gsmpc/accounts"
	//"github.com/deltaswapio/gsmpc/accounts/keystore"
	//"github.com/deltaswapio/gsmpc/common"
	//"github.com/deltaswapio/gsmpc/common/fdlimit"
	//"github.com/deltaswapio/gsmpc/consensus"
	//"github.com/deltaswapio/gsmpc/consensus/clique"
	//"github.com/deltaswapio/gsmpc/consensus/ethash"
	//"github.com/deltaswapio/gsmpc/core"
	//"github.com/deltaswapio/gsmpc/core/rawdb"
	//"github.com/deltaswapio/gsmpc/core/vm"
	//"github.com/deltaswapio/gsmpc/crypto"
	//"github.com/deltaswapio/gsmpc/eth"
	//"github.com/deltaswapio/gsmpc/eth/downloader"
	//"github.com/deltaswapio/gsmpc/eth/gasprice"
	//"github.com/deltaswapio/gsmpc/ethdb"
	//"github.com/deltaswapio/gsmpc/ethstats"
	//"github.com/deltaswapio/gsmpc/graphql"
	//"github.com/deltaswapio/gsmpc/internal/ethapi"
	//"github.com/deltaswapio/gsmpc/internal/flags"
	//"github.com/deltaswapio/gsmpc/les"
	//"github.com/deltaswapio/gsmpc/log"
	//"github.com/deltaswapio/gsmpc/metrics"
	//"github.com/deltaswapio/gsmpc/metrics/exp"
	//"github.com/deltaswapio/gsmpc/metrics/influxdb"
	//"github.com/deltaswapio/gsmpc/miner"
	//"github.com/deltaswapio/gsmpc/node"
	//"github.com/deltaswapio/gsmpc/p2p"
	//"github.com/deltaswapio/gsmpc/p2p/discv5"
	//"github.com/deltaswapio/gsmpc/p2p/enode"
	//"github.com/deltaswapio/gsmpc/p2p/nat"
	//"github.com/deltaswapio/gsmpc/p2p/netutil"
	//"github.com/deltaswapio/gsmpc/params"
	//pcsclite "github.com/gballet/go-libpcsclite"
	cli "gopkg.in/urfave/cli.v1"
)

// MigrateFlags sets the global flag from a local flag when it's set.
// This is a temporary function used for migrating old command/flags to the
// new format.
//
// e.g. geth account new --keystore /tmp/mykeystore --lightkdf
//
// is equivalent after calling this method with:
//
// geth --keystore /tmp/mykeystore --lightkdf account new
//
// This allows the use of the existing configuration functionality.
// When all flags are migrated this function can be removed and the existing
// configuration functionality must be changed that is uses local flags
func MigrateFlags(action func(ctx *cli.Context) error) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		for _, name := range ctx.FlagNames() {
			if ctx.IsSet(name) {
				err := ctx.GlobalSet(name, ctx.String(name))
				if err != nil {
					return err
				}
			}
		}
		return action(ctx)
	}
}
