// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { AnchoredTextDB } from '../anchoredtext-db'
import { AnchoredTextService } from '../anchoredtext.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { LinkDB } from '../link-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// AnchoredTextDetailComponent is initilizaed from different routes
// AnchoredTextDetailComponentState detail different cases 
enum AnchoredTextDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Link_TextAtArrowEnd_SET,
	CREATE_INSTANCE_WITH_ASSOCIATION_Link_TextAtArrowStart_SET,
}

@Component({
	selector: 'app-anchoredtext-detail',
	templateUrl: './anchoredtext-detail.component.html',
	styleUrls: ['./anchoredtext-detail.component.css'],
})
export class AnchoredTextDetailComponent implements OnInit {

	// insertion point for declarations

	// the AnchoredTextDB of interest
	anchoredtext: AnchoredTextDB = new AnchoredTextDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: AnchoredTextDetailComponentState = AnchoredTextDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private anchoredtextService: AnchoredTextService,
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
			this.state = AnchoredTextDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = AnchoredTextDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "TextAtArrowEnd":
						// console.log("AnchoredText" + " is instanciated with back pointer to instance " + this.id + " Link association TextAtArrowEnd")
						this.state = AnchoredTextDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Link_TextAtArrowEnd_SET
						break;
					case "TextAtArrowStart":
						// console.log("AnchoredText" + " is instanciated with back pointer to instance " + this.id + " Link association TextAtArrowStart")
						this.state = AnchoredTextDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Link_TextAtArrowStart_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getAnchoredText()

		// observable for changes in structs
		this.anchoredtextService.AnchoredTextServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getAnchoredText()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getAnchoredText(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case AnchoredTextDetailComponentState.CREATE_INSTANCE:
						this.anchoredtext = new (AnchoredTextDB)
						break;
					case AnchoredTextDetailComponentState.UPDATE_INSTANCE:
						let anchoredtext = frontRepo.AnchoredTexts.get(this.id)
						console.assert(anchoredtext != undefined, "missing anchoredtext with id:" + this.id)
						this.anchoredtext = anchoredtext!
						break;
					// insertion point for init of association field
					case AnchoredTextDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Link_TextAtArrowEnd_SET:
						this.anchoredtext = new (AnchoredTextDB)
						this.anchoredtext.Link_TextAtArrowEnd_reverse = frontRepo.Links.get(this.id)!
						break;
					case AnchoredTextDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Link_TextAtArrowStart_SET:
						this.anchoredtext = new (AnchoredTextDB)
						this.anchoredtext.Link_TextAtArrowStart_reverse = frontRepo.Links.get(this.id)!
						break;
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.anchoredtext.Link_TextAtArrowEnd_reverse != undefined) {
			if (this.anchoredtext.Link_TextAtArrowEndDBID == undefined) {
				this.anchoredtext.Link_TextAtArrowEndDBID = new NullInt64
			}
			this.anchoredtext.Link_TextAtArrowEndDBID.Int64 = this.anchoredtext.Link_TextAtArrowEnd_reverse.ID
			this.anchoredtext.Link_TextAtArrowEndDBID.Valid = true
			if (this.anchoredtext.Link_TextAtArrowEndDBID_Index == undefined) {
				this.anchoredtext.Link_TextAtArrowEndDBID_Index = new NullInt64
			}
			this.anchoredtext.Link_TextAtArrowEndDBID_Index.Valid = true
			this.anchoredtext.Link_TextAtArrowEnd_reverse = new LinkDB // very important, otherwise, circular JSON
		}
		if (this.anchoredtext.Link_TextAtArrowStart_reverse != undefined) {
			if (this.anchoredtext.Link_TextAtArrowStartDBID == undefined) {
				this.anchoredtext.Link_TextAtArrowStartDBID = new NullInt64
			}
			this.anchoredtext.Link_TextAtArrowStartDBID.Int64 = this.anchoredtext.Link_TextAtArrowStart_reverse.ID
			this.anchoredtext.Link_TextAtArrowStartDBID.Valid = true
			if (this.anchoredtext.Link_TextAtArrowStartDBID_Index == undefined) {
				this.anchoredtext.Link_TextAtArrowStartDBID_Index = new NullInt64
			}
			this.anchoredtext.Link_TextAtArrowStartDBID_Index.Valid = true
			this.anchoredtext.Link_TextAtArrowStart_reverse = new LinkDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case AnchoredTextDetailComponentState.UPDATE_INSTANCE:
				this.anchoredtextService.updateAnchoredText(this.anchoredtext, this.GONG__StackPath)
					.subscribe(anchoredtext => {
						this.anchoredtextService.AnchoredTextServiceChanged.next("update")
					});
				break;
			default:
				this.anchoredtextService.postAnchoredText(this.anchoredtext, this.GONG__StackPath).subscribe(anchoredtext => {
					this.anchoredtextService.AnchoredTextServiceChanged.next("post")
					this.anchoredtext = new (AnchoredTextDB) // reset fields
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

			dialogData.ID = this.anchoredtext.ID!
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
			dialogData.ID = this.anchoredtext.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "AnchoredText"
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
			ID: this.anchoredtext.ID,
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
		if (this.anchoredtext.Name == "") {
			this.anchoredtext.Name = event.value.Name
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
