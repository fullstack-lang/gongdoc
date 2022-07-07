import * as joint from 'jointjs';
import * as gongdoc from 'gongdoc'
import { GongdocCommandService } from 'gongdoc';
import { MatLabel } from '@angular/material/form-field';

export function newUmlNote(note: gongdoc.NoteDB,
    positionService: gongdoc.PositionService,
    gongdocCommandSingloton: gongdoc.GongdocCommandDB,
    gongdocCommandService: gongdoc.GongdocCommandService): joint.shapes.basic.Rect {

    // fetch the fields, it must belong to the current diagram
    // and the type must match the note type
    var attributes = new Array<string>()


    let noteTitle = note.Body

    var rect = new joint.shapes.standard.Rectangle(
        {
            position: {
                x: note.X,
                y: note.Y
            },
            size: { width: note.Width, height: note.Heigth },
            name: [noteTitle],
            methods: [],
            // store relevant attributes for working when callback are invoked
            note: note,
            positionService: positionService,
            gongdocCommandSingloton: gongdocCommandSingloton,
            gongdocCommandService: gongdocCommandService
        }
    )
    let width = noteTitle.length * 12
    let lines = noteTitle.split(/\r\n|\r|\n/)
    let maxLength = 0
    for (let lineNb = 0; lineNb < lines.length; lineNb++) {
        if (lines[lineNb].length > maxLength)  {
            maxLength = lines[lineNb].length
        }
    }

    rect.resize(
        300, noteTitle.split(/\r\n|\r|\n/).length * 18
    )
    rect.attr({
        body: {
            rx: 10, // add a corner radius
            ry: 10,
            fill: '#ADD8E6'
        },
        text: {
            text: noteTitle
        }
    })

    return rect

}