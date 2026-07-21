package cli

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseFlags_FilePath(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
	}()

	os.Args = []string{"checker", "config.yaml"}

	opts, err := ParseFlags()

	require.NoError(t, err)
	assert.False(t, opts.Silent)
	assert.False(t, opts.Stdin)
	assert.Equal(t, "config.yaml", opts.Path)
}

func TestParseFlags_SilentShort(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
	}()

	os.Args = []string{"checker", "-s", "config.yaml"}

	opts, err := ParseFlags()

	require.NoError(t, err)
	assert.True(t, opts.Silent)
	assert.False(t, opts.Stdin)
	assert.Equal(t, "config.yaml", opts.Path)
}

func TestParseFlags_SilentLong(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
	}()

	os.Args = []string{"checker", "--silent", "config.yaml"}

	opts, err := ParseFlags()

	require.NoError(t, err)
	assert.True(t, opts.Silent)
	assert.False(t, opts.Stdin)
	assert.Equal(t, "config.yaml", opts.Path)
}

func TestParseFlags_Stdin(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
	}()

	os.Args = []string{"checker", "--stdin"}

	opts, err := ParseFlags()

	require.NoError(t, err)
	assert.True(t, opts.Stdin)
	assert.False(t, opts.Silent)
	assert.Empty(t, opts.Path)
}

func TestParseFlags_NoPath(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
	}()

	os.Args = []string{"checker"}

	_, err := ParseFlags()

	require.Error(t, err)
	assert.Equal(t, "usage: checker <config-file>", err.Error())
}
