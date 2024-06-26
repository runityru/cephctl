package main

import (
	"context"
	"fmt"
	"os"

	kingpin "github.com/alecthomas/kingpin/v2"
	log "github.com/sirupsen/logrus"

	"github.com/runityru/cephctl/ceph"
	applyCmd "github.com/runityru/cephctl/commands/apply"
	diffCmd "github.com/runityru/cephctl/commands/diff"
	dumpCephConfigCmd "github.com/runityru/cephctl/commands/dump/cephconfig"
	dumpCephOSDConfigCmd "github.com/runityru/cephctl/commands/dump/cephosdconfig"
	healthcheckCmd "github.com/runityru/cephctl/commands/healthcheck"
	"github.com/runityru/cephctl/differ"
	"github.com/runityru/cephctl/printer"
	"github.com/runityru/cephctl/service"
)

var (
	appVersion     = "n/a (dev build)"
	buildTimestamp = "undefined"

	app = kingpin.New("cephctl", "Small utility to control Ceph cluster configuration just like any other declarative configuration")

	cephBinary = app.
			Flag("ceph-binary", "Specify path to ceph binary").
			Short('b').
			Envar("CEPHCTL_CEPH_BINARY").
			Default("/usr/bin/ceph").
			String()

	debug = app.
		Flag("debug", "Enable debug mode").
		Short('d').
		Envar("CEPHCTL_DEBUG").
		Bool()

	trace = app.
		Flag("trace", "Enable trace mode (debug mode on steroids)").
		Short('t').
		Envar("CEPHCTL_TRACE").
		Bool()

	colorize = app.
			Flag("color", "Colorize diff output").
			Short('c').
			Envar("CEPHCTL_COLOR").
			Default("true").
			Bool()

	apply         = app.Command("apply", "Apply ceph configuration")
	applySpecFile = apply.Arg("filename", "Filename with configuration specification").Required().String()

	diff = app.Command("diff", "Show difference between running and desired configurations")

	diffSpecFile = diff.Arg("filename", "Filename with configuration specification").Required().String()

	dump              = app.Command("dump", "Dump runtime configuration")
	dumpCephConfig    = dump.Command("cephconfig", "dump Ceph runtime configuration")
	dumpCephOSDConfig = dump.Command("cephosdconfig", "dump Ceph OSD configuration")

	healthcheck = app.Command("healthcheck", "Perform a cluster healthcheck and print report")

	version = app.Command("version", "Print version and exit")
)

func main() {
	ctx := context.Background()
	appCmd := kingpin.MustParse(app.Parse(os.Args[1:]))

	if *trace {
		log.SetLevel(log.TraceLevel)
		log.SetFormatter(&log.TextFormatter{
			FullTimestamp: true,
		})
		log.Trace("Trace mode is enabled. Beware of verbosity!")
	} else if *debug {
		log.SetLevel(log.DebugLevel)
		log.SetFormatter(&log.TextFormatter{
			FullTimestamp: true,
		})
		log.Debug("Debug mode is enabled.")
	}

	svc := service.New(ceph.New(*cephBinary), differ.New())
	prntr := printer.New(*colorize)

	switch appCmd {
	case apply.FullCommand():
		log.Debug("running apply command")
		if err := applyCmd.Apply(ctx, applyCmd.ApplyConfig{
			Service:  svc,
			SpecFile: *applySpecFile,
		}); err != nil {
			panic(err)
		}

	case diff.FullCommand():
		log.Debug("running diff command")
		if err := diffCmd.Diff(ctx, diffCmd.DiffConfig{
			Printer:  prntr,
			Service:  svc,
			SpecFile: *diffSpecFile,
		}); err != nil {
			panic(err)
		}

	case dumpCephConfig.FullCommand():
		log.Debug("running dump cephconfig command")
		if err := dumpCephConfigCmd.DumpCephConfig(ctx, dumpCephConfigCmd.DumpCephConfigConfig{
			Printer: prntr,
			Service: svc,
		}); err != nil {
			panic(err)
		}

	case dumpCephOSDConfig.FullCommand():
		log.Debug("running dump cephosdconfig command")
		if err := dumpCephOSDConfigCmd.DumpCephOSDConfig(ctx, dumpCephOSDConfigCmd.DumpCephOSDConfigConfig{
			Printer: prntr,
			Service: svc,
		}); err != nil {
			panic(err)
		}

	case healthcheck.FullCommand():
		if err := healthcheckCmd.Healthcheck(ctx, healthcheckCmd.HealthcheckConfig{
			Printer: prntr,
			Service: svc,
		}); err != nil {
			panic(err)
		}

	case version.FullCommand():
		fmt.Printf(
			"%s %s built @ %s\n",
			os.Args[0], appVersion, buildTimestamp,
		)
		os.Exit(1)
	}
}
