package diff

import (
	"github.com/rs/zerolog/log"
	"github.com/sergi/go-diff/diffmatchpatch"
)

type EditScript string;

type Differ struct {
	diff_matcher *diffmatchpatch.DiffMatchPatch
	ToStd        chan string
	FromStd      chan string
	ToClient     chan EditScript
	FromClient   chan EditScript
}

func NewDiffer() *Differ {
	return &Differ{diff_matcher: diffmatchpatch.New(), ToStd: make(chan string),
		FromStd: make(chan string), ToClient: make(chan EditScript), FromClient: make(chan EditScript)}
}

func (d *Differ) StartDiffer(fmsg string) {
	var prev string = fmsg
    log.Debug().Msgf("StartDiffer:First Message :%s", fmsg);

	go func() {
		//Takes input from the stdin diffs it with the prev string and sends to *ToClient chan*
		for txt := range d.FromStd {
			edit_script := d.toDiff(prev, txt)
			d.ToClient <- edit_script
		}
	}()

	for edit_script := range d.FromClient {
		updated := d.FromDiff(prev, edit_script)
		d.ToStd <- updated
		prev = updated
	}
}

func (d *Differ) toDiff(src, dst string) EditScript {
	diffs := d.diff_matcher.DiffMain(src, dst, false)
	return EditScript(d.diff_matcher.DiffToDelta(diffs))
}

func (d *Differ) FromDiff(src string, edit_script EditScript) string {
	differ := d.diff_matcher
	dst, err := differ.DiffFromDelta(src, string(edit_script))

	if err != nil {
        log.Fatal().Err(err)
	}

    updated := differ.DiffText2(dst) ;
    log.Debug().Msgf("FromDiff: %s and %s => %s\n",src, edit_script, updated);
	return updated
}
