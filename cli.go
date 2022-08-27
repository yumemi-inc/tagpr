package rcpr

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
)

const cmdName = "rcpr"

func printVersion(out io.Writer) error {
	_, err := fmt.Fprintf(out, "%s v%s (rev:%s)\n", cmdName, version, revision)
	return err
}

// Run the rcpr
func Run(ctx context.Context, argv []string, outStream, errStream io.Writer) error {
	log.SetOutput(errStream)
	fs := flag.NewFlagSet(
		fmt.Sprintf("%s (v%s rev:%s)", cmdName, version, revision), flag.ContinueOnError)
	fs.SetOutput(errStream)
	ver := fs.Bool("version", false, "display version")
	if err := fs.Parse(argv); err != nil {
		return err
	}
	if *ver {
		return printVersion(outStream)
	}

	rp, err := newRcpr(ctx, &commander{
		gitPath: "git", outStream: outStream, errStream: errStream, dir: "."})
	if err != nil {
		return err
	}
	return rp.Run(ctx)
}