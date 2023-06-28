package diff

import (
	"github.com/sergi/go-diff/diffmatchpatch"
)

func Differ(src, dst string) string {
	differ := diffmatchpatch.New()
	return differ.DiffToDelta(differ.DiffMain(src, dst, false))
}

func DeDiffer(src, edit_script string) string {
	differ := diffmatchpatch.New()
	dst, err := differ.DiffFromDelta(src, edit_script)

	if err != nil {
	}

	return differ.DiffText2(dst)
}
