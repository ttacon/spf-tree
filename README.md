# spf-tree
Shows the SPF lookups for a mail host as a tree. Currently, this only looks at
`include:`d lookups.

## Installation

Only installation from source is supported at this time. So as long as you have
Go installed (v1.15+), you can do:

```sh
go get github.com/ttacon/spf-tree/cmd/spf-tree
```

## Usage

Provide `spf-tree` a host and see that the SPF lookups look like:

```sh
$ spf-tree -h=asana.com
.
├── _spf.google.com
│   ├── _netblocks.google.com
│   ├── _netblocks2.google.com
│   └── _netblocks3.google.com
├── mktomail.com
└── mg-spf.greenhouse.io

$ spf-tree -h=duo.com
.
├── _spf.google.com
│   ├── _netblocks.google.com
│   ├── _netblocks2.google.com
│   └── _netblocks3.google.com
├── _spf.salesforce.com
├── mktomail.com
├── stspg-customer.com
├── _spf.intacct.com
└── servers.mcsv.net
```
