import * as gongdoc from 'gongdoc'

export function informBackEndOfSelection(cellView: joint.dia.CellView) {

    let umlClassShape = cellView.model

    let gongStructShape = umlClassShape.attributes['classshape'] as gongdoc.GongShapeDB
    let gongStructShapeService = umlClassShape.attributes['gongStructShapeService'] as gongdoc.GongShapeService

    // if selected object is not a classshape, move on
    if (gongStructShape == undefined) {
        return
    }

    gongStructShape.IsSelected = true
    gongStructShapeService.updateGongShape(gongStructShape, "").subscribe(
        classhape => {
            console.log("gongStructShape updated")
        }
    )
}