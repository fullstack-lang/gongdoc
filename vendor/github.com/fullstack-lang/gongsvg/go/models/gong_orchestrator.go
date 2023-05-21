package models

// insertion point
// AnchoredTextOrchestrator
type AnchoredTextOrchestrator struct {
}

func (orchestrator *AnchoredTextOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedAnchoredText, backRepoAnchoredText *AnchoredText) {

	stagedAnchoredText.OnAfterUpdate(gongsvgStage, stagedAnchoredText, backRepoAnchoredText)
}
// LineOrchestrator
type LineOrchestrator struct {
}

func (orchestrator *LineOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedLine, backRepoLine *Line) {

	stagedLine.OnAfterUpdate(gongsvgStage, stagedLine, backRepoLine)
}
// LinkOrchestrator
type LinkOrchestrator struct {
}

func (orchestrator *LinkOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedLink, backRepoLink *Link) {

	stagedLink.OnAfterUpdate(gongsvgStage, stagedLink, backRepoLink)
}
// RectOrchestrator
type RectOrchestrator struct {
}

func (orchestrator *RectOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedRect, backRepoRect *Rect) {

	stagedRect.OnAfterUpdate(gongsvgStage, stagedRect, backRepoRect)
}
// SVGOrchestrator
type SVGOrchestrator struct {
}

func (orchestrator *SVGOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedSVG, backRepoSVG *SVG) {

	stagedSVG.OnAfterUpdate(gongsvgStage, stagedSVG, backRepoSVG)
}

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *StageStruct) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case AnchoredText:
		stage.OnAfterAnchoredTextUpdateCallback = new(AnchoredTextOrchestrator)
	case Line:
		stage.OnAfterLineUpdateCallback = new(LineOrchestrator)
	case Link:
		stage.OnAfterLinkUpdateCallback = new(LinkOrchestrator)
	case Rect:
		stage.OnAfterRectUpdateCallback = new(RectOrchestrator)
	case SVG:
		stage.OnAfterSVGUpdateCallback = new(SVGOrchestrator)

	}

}
