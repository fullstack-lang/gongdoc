// insertion point for imports
import { VerticeDB } from './vertice-db'
import { GongShapeDB } from './gongshape-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LinkDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Identifier: string = ""
	Fieldtypename: string = ""
	TargetMultiplicity: string = ""
	SourceMultiplicity: string = ""

	// insertion point for other declarations
	Middlevertice?: VerticeDB
	MiddleverticeID: NullInt64 = new NullInt64 // if pointer is null, Middlevertice.ID = 0

	GongShape_LinksDBID: NullInt64 = new NullInt64
	GongShape_LinksDBID_Index: NullInt64  = new NullInt64 // store the index of the link instance in GongShape.Links
	GongShape_Links_reverse?: GongShapeDB 

}
