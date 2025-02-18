package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var create_configmapCmd = &cobra.Command{
	Use:     "configmap NAME [--from-file=[key=]source] [--from-literal=key1=value1] [--dry-run=server|client|none]",
	Short:   "Create a config map from a local file, directory or literal value",
	Aliases: []string{"cm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_configmapCmd).Standalone()

	create_configmapCmd.Flags().Bool("allow-missing-template-keys", true, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	create_configmapCmd.Flags().Bool("append-hash", false, "Append a hash of the configmap to its name.")
	create_configmapCmd.Flags().String("dry-run", "none", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	create_configmapCmd.Flags().String("field-manager", "kubectl-create", "Name of the manager used to track field ownership.")
	create_configmapCmd.Flags().StringSlice("from-env-file", []string{}, "Specify the path to a file to read lines of key=val pairs to create a configmap.")
	create_configmapCmd.Flags().StringSlice("from-file", []string{}, "Key file can be specified using its file path, in which case file basename will be used as configmap key, or optionally with a key and file path, in which case the given key will be used.  Specifying a directory will iterate each named file in the directory whose basename is a valid configmap key.")
	create_configmapCmd.Flags().StringArray("from-literal", []string{}, "Specify a key and literal value to insert in configmap (i.e. mykey=somevalue)")
	create_configmapCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	create_configmapCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.")
	create_configmapCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	create_configmapCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	create_configmapCmd.Flags().String("validate", "strict", "Must be one of: strict (or true), warn, ignore (or false).")
	create_configmapCmd.Flag("dry-run").NoOptDefVal = "unchanged"
	create_configmapCmd.Flag("validate").NoOptDefVal = "strict"
	createCmd.AddCommand(create_configmapCmd)

	carapace.Gen(create_configmapCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":       kubectl.ActionDryRunModes(),
		"from-env-file": carapace.ActionFiles(),
		"output":        kubectl.ActionOutputFormats(),
		"template":      carapace.ActionFiles(),
		"validate":      kubectl.ActionValidationModes(),
	})
}
