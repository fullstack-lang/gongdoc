import * as gongdoc from 'gongdoc'

export function informBackEndOfSelection(cellView: joint.dia.CellView) {

    let umlClassShape = cellView.model

    let classhape = umlClassShape.attributes['classshape'] as gongdoc.ClassshapeDB
    let classshapeService = umlClassShape.attributes['classshapeService'] as gongdoc.ClassshapeService

    // if selected object is not a classshape, move on
    if (classhape == undefined) {
        return
    }

    classhape.IsSelected = true
    classshapeService.updateClassshape(classhape).subscribe(
        classhape => {
            console.log("classhape updated")
        }
    )
}