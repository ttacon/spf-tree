package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strings"

	"github.com/xlab/treeprint"
)

var (
	host = flag.String("h", "", "host to print the SPF tree for")

	hostMatcher = regexp.MustCompile(`^\w+\.\w+\.\w+$`)
)

func main() {
	flag.Parse()

	if len(*host) == 0 {
		log.Println("must provide a host")
		os.Exit(1)
	}

	tree := treeprint.New()
	if err := nodesForHost(*host, tree); err != nil {
		log.Println("failed to finish SPF resolution, err: ", err)
		os.Exit(1)
	}

	fmt.Println(tree.String())
}

// nodesForHost adds all subnodes to the given tree recursively.
// If there is an error encountered while retrieving the DNS TXT records for
// the given host, that error will be returned.
func nodesForHost(host string, tree treeprint.Tree) error {
	records, err := net.LookupTXT(host)
	if err != nil {
		return err
	}

	var spfRecord string
	for _, record := range records {
		if strings.HasPrefix(record, "v=spf1") {
			spfRecord = record
			break
		}
	}
	if len(spfRecord) == 0 {
		return nil
	}

	pieces := strings.Split(spfRecord, " ")
	for _, piece := range pieces {
		cleaned := strings.TrimSpace(piece)
		if !strings.HasPrefix(cleaned, "include:") {
			continue
		}

		stripped := strings.TrimPrefix(cleaned, "include:")
		node := tree.AddBranch(stripped)
		if addr := net.ParseIP(host); addr != nil {
			// It's a CIDR range, don't dig deeper.
			continue
		}

		// It's a host and not a CIDR block, go deeper.
		if err := nodesForHost(stripped, node); err != nil {
			return err
		}
	}

	return nil
}
