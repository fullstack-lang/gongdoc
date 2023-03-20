// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { GongShapeDB } from '../gongshape-db'
import { GongShapeService } from '../gongshape.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { ClassdiagramDB } from '../classdiagram-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// GongShapeDetailComponent is initilizaed from different routes
// GongShapeDetailComponentState detail different cases 
enum GongShapeDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Classdiagram_GongStructShapes_SET,
}

@Component({
	selector: 'app-gongshape-detail',
	templateUrl: './gongshape-detail.component.html',
	styleUrls: ['./gongshape-detail.component.css'],
})
export class GongShapeDetailComponent implements OnInit {

	// insertion point for declarations
	ShowNbInstancesFormControl: UntypedFormControl = new UntypedFormControl(false);
	IsSelectedFormControl: UntypedFormControl = new UntypedFormControl(false);

	// the GongShapeDB of interest
	gongshape: GongShapeDB = new GongShapeDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: GongShapeDetailComponentState = GongShapeDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private gongshapeService: GongShapeService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private activatedRoute: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {
		this.GONG__StackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')!;

		this.activatedRoute.params.subscribe(params => {
			this.onChangedActivatedRoute()
		});
	}
	onChangedActivatedRoute(): void {

		// compute state
		this.id = +this.activatedRoute.snapshot.paramMap.get('id')!;
		this.originStruct = this.activatedRoute.snapshot.paramMap.get('originStruct')!;
		this.originStructFieldName = this.activatedRoute.snapshot.paramMap.get('originStructFieldName')!;

		this.GONG__StackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')!;

		const association = this.activatedRoute.snapshot.paramMap.get('association');
		if (this.id == 0) {
			this.state = GongShapeDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = GongShapeDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "GongStructShapes":
						// console.log("GongShape" + " is instanciated with back pointer to instance " + this.id + " Classdiagram association GongStructShapes")
						this.state = GongShapeDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Classdiagram_GongStructShapes_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getGongShape()

		// observable for changes in structs
		this.gongshapeService.GongShapeServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getGongShape()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getGongShape(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case GongShapeDetailComponentState.CREATE_INSTANCE:
						this.gongshape = new (GongShapeDB)
						break;
					case GongShapeDetailComponentState.UPDATE_INSTANCE:
						let gongshape = frontRepo.GongShapes.get(this.id)
						console.assert(gongshape != undefined, "missing gongshape with id:" + this.id)
						this.gongshape = gongshape!
						break;
					// insertion point for init of association field
					case GongShapeDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Classdiagram_GongStructShapes_SET:
						this.gongshape = new (GongShapeDB)
						this.gongshape.Classdiagram_GongStructShapes_reverse = frontRepo.Classdiagrams.get(this.id)!
						break;
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
				this.ShowNbInstancesFormControl.setValue(this.gongshape.ShowNbInstances)
				this.IsSelectedFormControl.setValue(this.gongshape.IsSelected)
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field
		if (this.gongshape.PositionID == undefined) {
			this.gongshape.PositionID = new NullInt64
		}
		if (this.gongshape.Position != undefined) {
			this.gongshape.PositionID.Int64 = this.gongshape.Position.ID
			this.gongshape.PositionID.Valid = true
		} else {
			this.gongshape.PositionID.Int64 = 0
			this.gongshape.PositionID.Valid = true
		}
		this.gongshape.ShowNbInstances = this.ShowNbInstancesFormControl.value
		this.gongshape.IsSelected = this.IsSelectedFormControl.value

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.gongshape.Classdiagram_GongStructShapes_reverse != undefined) {
			if (this.gongshape.Classdiagram_GongStructShapesDBID == undefined) {
				this.gongshape.Classdiagram_GongStructShapesDBID = new NullInt64
			}
			this.gongshape.Classdiagram_GongStructShapesDBID.Int64 = this.gongshape.Classdiagram_GongStructShapes_reverse.ID
			this.gongshape.Classdiagram_GongStructShapesDBID.Valid = true
			if (this.gongshape.Classdiagram_GongStructShapesDBID_Index == undefined) {
				this.gongshape.Classdiagram_GongStructShapesDBID_Index = new NullInt64
			}
			this.gongshape.Classdiagram_GongStructShapesDBID_Index.Valid = true
			this.gongshape.Classdiagram_GongStructShapes_reverse = new ClassdiagramDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case GongShapeDetailComponentState.UPDATE_INSTANCE:
				this.gongshapeService.updateGongShape(this.gongshape, this.GONG__StackPath)
					.subscribe(gongshape => {
						this.gongshapeService.GongShapeServiceChanged.next("update")
					});
				break;
			default:
				this.gongshapeService.postGongShape(this.gongshape, this.GONG__StackPath).subscribe(gongshape => {
					this.gongshapeService.GongShapeServiceChanged.next("post")
					this.gongshape = new (GongShapeDB) // reset fields
				});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string, selectionMode: string,
		sourceField: string, intermediateStructField: string, nextAssociatedStruct: string) {

		console.log("mode " + selectionMode)

		const dialogConfig = new MatDialogConfig();

		let dialogData = new DialogData();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.width = "50%"
		dialogConfig.height = "50%"
		if (selectionMode == SelectionMode.ONE_MANY_ASSOCIATION_MODE) {

			dialogData.ID = this.gongshape.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(AssociatedStruct).get(
					AssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}
		if (selectionMode == SelectionMode.MANY_MANY_ASSOCIATION_MODE) {
			dialogData.ID = this.gongshape.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "GongShape"
			dialogData.SourceField = sourceField

			// set up the intermediate struct
			dialogData.IntermediateStruct = AssociatedStruct
			dialogData.IntermediateStructField = intermediateStructField

			// set up the end struct
			dialogData.NextAssociationStruct = nextAssociatedStruct

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(nextAssociatedStruct).get(
					nextAssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}

	}

	openDragAndDropOrdering(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.data = {
			ID: this.gongshape.ID,
			ReversePointer: reverseField,
			OrderingMode: true,
			GONG__StackPath: this.GONG__StackPath,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfSortingComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'SortingComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
		});
	}

	fillUpNameIfEmpty(event: { value: { Name: string; }; }) {
		if (this.gongshape.Name == "") {
			this.gongshape.Name = event.value.Name
		}
	}

	toggleTextArea(fieldName: string) {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			let displayAsTextArea = this.mapFields_displayAsTextArea.get(fieldName)
			this.mapFields_displayAsTextArea.set(fieldName, !displayAsTextArea)
		} else {
			this.mapFields_displayAsTextArea.set(fieldName, true)
		}
	}

	isATextArea(fieldName: string): boolean {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			return this.mapFields_displayAsTextArea.get(fieldName)!
		} else {
			return false
		}
	}

	compareObjects(o1: any, o2: any) {
		if (o1?.ID == o2?.ID) {
			return true;
		}
		else {
			return false
		}
	}
}
