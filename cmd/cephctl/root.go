package main

import (
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

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

// Variables are bound to the root command's persistent flags via BoolVar/
// StringVarP. Their values are updated when cmd.Flags().Set(...) is called
// (including from env).
var (
	cephBinary string
	debug      bool
	trace      bool
	colorize   bool
)

var rootCmd = &cobra.Command{
	Use:   "cephctl",
	Short: "Small utility to control Ceph cluster configuration just like any other declarative configuration",
	// SilenceUsage: do not print usage when a RunE handler returns an error.
	SilenceUsage: true,
	// SilenceErrors is left false (default): cobra prints the error text to
	// stderr, and main() terminates the process with code 1.
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Env support: an environment variable value is applied as the default.
		// If the flag was set explicitly by the user (f.Changed), env is ignored.
		// This mirrors kingpin's Envar semantics (env = default, flag overrides).
		applyStringEnv(cmd, "ceph-binary", "CEPHCTL_CEPH_BINARY")
		applyBoolEnv(cmd, "debug", "CEPHCTL_DEBUG")
		applyBoolEnv(cmd, "trace", "CEPHCTL_TRACE")
		applyBoolEnv(cmd, "color", "CEPHCTL_COLOR")

		if trace {
			log.SetLevel(log.TraceLevel)
			log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
			log.Trace("Trace mode is enabled. Beware of verbosity!")
		} else if debug {
			log.SetLevel(log.DebugLevel)
			log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
			log.Debug("Debug mode is enabled.")
		}
	},
	// Without a subcommand, print the root command's help.
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	// Global (persistent) flags — semantics preserved from the kingpin version.
	rootCmd.PersistentFlags().StringVarP(&cephBinary, "ceph-binary", "b", "/usr/bin/ceph", "Specify path to ceph binary")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug mode")
	rootCmd.PersistentFlags().BoolVarP(&trace, "trace", "t", false, "Enable trace mode (debug mode on steroids)")
	rootCmd.PersistentFlags().BoolVarP(&colorize, "color", "c", true, "Colorize diff output")

	rootCmd.AddCommand(applyCommand)
	rootCmd.AddCommand(diffCommand)
	rootCmd.AddCommand(dumpCommand)
	rootCmd.AddCommand(healthcheckCommand)
	rootCmd.AddCommand(versionCommand)

	dumpCommand.AddCommand(dumpCephConfigCommand)
	dumpCommand.AddCommand(dumpCephOSDConfigCommand)
}

var applyCommand = &cobra.Command{
	Use:   "apply <filename>",
	Short: "Apply ceph configuration",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Debug("running apply command")
		svc := service.New(ceph.New(cephBinary), differ.New())
		return applyCmd.Apply(cmd.Context(), applyCmd.ApplyConfig{
			Service:  svc,
			SpecFile: args[0],
		})
	},
}

var diffCommand = &cobra.Command{
	Use:   "diff <filename>",
	Short: "Show difference between running and desired configurations",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Debug("running diff command")
		svc := service.New(ceph.New(cephBinary), differ.New())
		prntr := printer.New(colorize)
		return diffCmd.Diff(cmd.Context(), diffCmd.DiffConfig{
			Printer:  prntr,
			Service:  svc,
			SpecFile: args[0],
		})
	},
}

var dumpCommand = &cobra.Command{
	Use:   "dump",
	Short: "Dump runtime configuration",
	// dump without a child command — print the group's help.
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var dumpCephConfigCommand = &cobra.Command{
	Use:   "cephconfig",
	Short: "dump Ceph runtime configuration",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Debug("running dump cephconfig command")
		svc := service.New(ceph.New(cephBinary), differ.New())
		prntr := printer.New(colorize)
		return dumpCephConfigCmd.DumpCephConfig(cmd.Context(), dumpCephConfigCmd.DumpCephConfigConfig{
			Printer: prntr,
			Service: svc,
		})
	},
}

var dumpCephOSDConfigCommand = &cobra.Command{
	Use:   "cephosdconfig",
	Short: "dump Ceph OSD configuration",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Debug("running dump cephosdconfig command")
		svc := service.New(ceph.New(cephBinary), differ.New())
		prntr := printer.New(colorize)
		return dumpCephOSDConfigCmd.DumpCephOSDConfig(cmd.Context(), dumpCephOSDConfigCmd.DumpCephOSDConfigConfig{
			Printer: prntr,
			Service: svc,
		})
	},
}

var healthcheckCommand = &cobra.Command{
	Use:   "healthcheck",
	Short: "Perform a cluster healthcheck and print report",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		svc := service.New(ceph.New(cephBinary), differ.New())
		prntr := printer.New(colorize)
		return healthcheckCmd.Healthcheck(cmd.Context(), healthcheckCmd.HealthcheckConfig{
			Printer: prntr,
			Service: svc,
		})
	},
}

// applyStringEnv applies a string env value as the default for a flag, unless
// the flag was set explicitly by the user.
func applyStringEnv(cmd *cobra.Command, flagName, envName string) {
	f := cmd.Flags().Lookup(flagName)
	if f == nil || f.Changed {
		return
	}
	if v, ok := os.LookupEnv(envName); ok {
		_ = cmd.Flags().Set(flagName, v)
	}
}

// applyBoolEnv applies a truthy env value (1/true/yes/on) as the default for a
// bool flag, unless the flag was set explicitly by the user.
func applyBoolEnv(cmd *cobra.Command, flagName, envName string) {
	f := cmd.Flags().Lookup(flagName)
	if f == nil || f.Changed {
		return
	}
	if v, ok := os.LookupEnv(envName); ok {
		_ = cmd.Flags().Set(flagName, strconv.FormatBool(parseBool(v)))
	}
}

// parseBool treats an env string as truthy for bool flags.
func parseBool(v string) bool {
	switch strings.ToLower(strings.TrimSpace(v)) {
	case "1", "true", "yes", "on":
		return true
	default:
		return false
	}
}
