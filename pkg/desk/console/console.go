package console

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/vitaliy-ukiru/find-pair-game/domain/entity"
	"github.com/vitaliy-ukiru/find-pair-game/domain/game"
	"github.com/vitaliy-ukiru/find-pair-game/pkg/desk/console/consolerender"
	"github.com/vitaliy-ukiru/find-pair-game/pkg/visual"
)

type Desk struct {
	g      *game.Game
	reader *bufio.Reader

	renderer       *visual.DeskRender
	texts          Texts
	wrongCellSleep time.Duration
	startWithZero  bool
	finish         bool
}

func New(
	g *game.Game,
	reader *bufio.Reader,
	cardPlacer visual.CardPlacer,
	cardRenderer visual.CardRenderer,
	tableRenderer *consolerender.TableRenderer,
	texts Texts,
) *Desk {
	return &Desk{
		g:      g,
		reader: reader,
		renderer: visual.NewDeskRender(
			g.Sizes(),
			cardRenderer,
			tableRenderer,
			cardPlacer,
		),
		texts:          texts,
		wrongCellSleep: 3 * time.Second,
	}
}

func (d *Desk) SetStartWithZero(startWithZero bool) {
	d.startWithZero = startWithZero
}

func (d *Desk) SetWrongCellSleep(wrongCellSleep time.Duration) {
	d.wrongCellSleep = wrongCellSleep
}

type Texts struct {
	GuessAll             string
	InputPrompt          string
	WrongStep            string
	ErrorOutOfBoard      string
	ErrorIncorrectFormat string
	RenderError          string
	PrefixError          string
}

func (t Texts) withError(errMsg string) string {
	return t.PrefixError + errMsg
}

type nextCallbackFn func() nextCallbackFn

func (d *Desk) onErrorFrame(err error) nextCallbackFn {
	return func() nextCallbackFn {
		msg := err.Error()
		if errors.Is(err, game.ErrOutOfRange) {
			msg = d.texts.ErrorOutOfBoard
		}

		d.renderMessage(d.texts.withError(msg))
		return d.processPlay()
	}
}

func (d *Desk) onFrame() nextCallbackFn {
	d.renderDefault()
	return d.processPlay()
}

func (d *Desk) processPlay() nextCallbackFn {
	point, err := d.readInput()
	if err != nil {
		if errors.Is(err, io.EOF) {
			os.Exit(0)
		}
		panic(err)
	}

	result, err := d.g.MakeClick(point)
	if err != nil {
		return d.onErrorFrame(err)
	}

	if result.Result == game.Wrong {
		d.processWrong(result)
	}
	switch result.Result {
	case game.Wrong:
		d.processWrong(result)
	case game.Finish:
		d.finish = true
	}
	return d.onFrame
}

func (d *Desk) Play() {
	fn := d.onFrame
	for d.isPlaying() {
		fn = fn()
		if fn == nil {
			fmt.Println("Next step is nil :/")
			os.Exit(-1)
		}
	}
	d.renderMessage(d.texts.GuessAll)
}

func (d *Desk) processWrong(result game.ClickResult) {
	//items[result.CurrentPoint] = game.BoardItem{
	//	Point:  result.CurrentPoint,
	//	Card:   result.Current,
	//	Status: game.CardOpen,
	//}
	//d.render(items, true, d.texts.WrongStep)
	d.render(
		true,
		d.texts.WrongStep,
		d.renderer.WithCardPlacer(visual.NewWrongCardPlacer(
			result.Current,
			*result.Prev,
			d.renderer.CardPlacer(),
		)),
	)
	time.Sleep(d.wrongCellSleep)
}

func (d *Desk) readInput() (entity.Point, error) {
	fmt.Print(d.texts.InputPrompt)
	var p entity.Point
	input, _, err := d.reader.ReadLine()
	if err != nil {
		return p, err
	}
	if _, err := fmt.Sscanf(string(input), "%d %d", &p.X, &p.Y); err != nil {
		d.renderMessage(d.texts.withError(d.texts.ErrorIncorrectFormat))
		return d.readInput()
	}
	if !d.startWithZero {
		p.X--
		p.Y--
	}
	return p, nil
}

func (d *Desk) isPlaying() bool {
	return !d.finish
}
