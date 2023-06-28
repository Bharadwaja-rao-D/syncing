package diff

import (
	"github.com/sergi/go-diff/diffmatchpatch"
	"log"
)

type Differ struct {
	diff_matcher *diffmatchpatch.DiffMatchPatch
	to_ws        chan string
	from_ws      chan string
}

func (d *Differ) ToDiff(src, dst string) string {
	diffs := d.diff_matcher.DiffMain(src, dst, false)
	return d.diff_matcher.DiffToDelta(diffs)
}

func (d *Differ) FromDiff(src, edit_script string) string {
	differ := d.diff_matcher
	dst, err := differ.DiffFromDelta(src, edit_script)

	if err != nil {
		log.Fatal(err)
	}

	return differ.DiffText2(dst)
}
