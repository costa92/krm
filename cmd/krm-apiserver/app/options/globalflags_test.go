package options

import (
	"flag"
	"github.com/spf13/pflag"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/cli/globalflag"
	"k8s.io/component-base/logs"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestAddCustomGlobalFlags(t *testing.T) {
	namedFlagSets := &cliflag.NamedFlagSets{}
	// flags.CommandLine.
	nfs := namedFlagSets.FlagSet("test")

	nfs.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)
	globalflag.AddGlobalFlags(nfs, "test-cmd")
	AddCustomGlobalFlags(nfs)

	actualFlag := []string{}
	nfs.VisitAll(func(flag *pflag.Flag) {
		actualFlag = append(actualFlag, flag.Name)
	})

	// Get all flags from flags.CommandLine, except flag `test.*`.
	wantedFlag := []string{"help"}
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	logs.AddFlags(pflag.CommandLine)
	normalizeFunc := nfs.GetNormalizeFunc()
	pflag.VisitAll(func(flag *pflag.Flag) {
		if !strings.Contains(flag.Name, "test.") {
			wantedFlag = append(wantedFlag, string(normalizeFunc(nfs, flag.Name)))
		}
	})
	sort.Strings(wantedFlag)

	if !reflect.DeepEqual(wantedFlag, actualFlag) {
		t.Errorf("[Default]: expected %+v, got %+v", wantedFlag, actualFlag)
	}
}
