import * as joint from 'jointjs';
import * as gongdoc from 'gongdoc'
import { GongdocCommandService } from 'gongdoc';

export function newUmlNote(note: gongdoc.NoteDB,
    positionService: gongdoc.PositionService,
    gongdocCommandSingloton: gongdoc.GongdocCommandDB,
    gongdocCommandService: gongdoc.GongdocCommandService): joint.shapes.basic.Rect {

    // fetch the fields, it must belong to the current diagram
    // and the type must match the note type
    var attributes = new Array<string>()


    let noteTitle = note.Body

    return new joint.shapes.basic.Rect(
        {
            position: {
                x: note.X,
                y: note.Y
            },
            size: { width: note.Width, height: note.Heigth },
            name: [noteTitle],
            methods: [],
            attrs: {
                '.uml-class-name-rect': {
                    fill: '#ff8450',
                    stroke: '#fff',
                    'stroke-width': 0.5,
                },
                '.uml-class-name-text': {
                    'font-family': 'Roboto'
                },
                '.uml-class-attrs-rect': {
                    fill: '#fe976a',
                    stroke: '#fff',
                    height: 10,
                    'stroke-width': 0.5,
                    'font-family': 'Roboto'
                },
                '.uml-class-methods-rect': {
                    fill: '#fe976a',
                    stroke: '#fff',
                    height: 0,
                    'stroke-width': 0
                },
                '.uml-class-attrs-text': {
                    'ref-y': 0,
                    'y-alignment': 'top',
                    'font-family': 'Roboto'
                }
            },

            // store relevant attributes for working when callback are invoked
            note: note,
            positionService: positionService,
            gongdocCommandSingloton: gongdocCommandSingloton,
            gongdocCommandService: gongdocCommandService
        }
    )

}