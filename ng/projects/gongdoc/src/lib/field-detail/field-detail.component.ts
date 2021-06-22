// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { FieldDB } from '../field-db'
import { FieldService } from '../field.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-field-detail',
	templateUrl: './field-detail.component.html',
	styleUrls: ['./field-detail.component.css'],
})
export class FieldDetailComponent implements OnInit {

	// insertion point for declarations

	// the FieldDB of interest
	field: FieldDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private fieldService: FieldService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {
		this.getField()

		// observable for changes in structs
		this.fieldService.FieldServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getField()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getField(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				if (id != 0 && association == undefined) {
					this.field = frontRepo.Fields.get(id)
				} else {
					this.field = new (FieldDB)
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization
		
		// insertion point for translation/nullation of each field
		
		// save from the front pointer space to the non pointer space for serialization
		if (association == undefined) {
			// insertion point for translation/nullation of each pointers
			if (this.field.Classshape_Fields_reverse != undefined) {
				if (this.field.Classshape_FieldsDBID == undefined) {
					this.field.Classshape_FieldsDBID = new NullInt64
				}
				this.field.Classshape_FieldsDBID.Int64 = this.field.Classshape_Fields_reverse.ID
				this.field.Classshape_FieldsDBID.Valid = true
				if (this.field.Classshape_FieldsDBID_Index == undefined) {
					this.field.Classshape_FieldsDBID_Index = new NullInt64
				}
				this.field.Classshape_FieldsDBID_Index.Valid = true
				this.field.Classshape_Fields_reverse = undefined // very important, otherwise, circular JSON
			}
		}

		if (id != 0 && association == undefined) {

			this.fieldService.updateField(this.field)
				.subscribe(field => {
					this.fieldService.FieldServiceChanged.next("update")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "Classshape_Fields":
					this.field.Classshape_FieldsDBID = new NullInt64
					this.field.Classshape_FieldsDBID.Int64 = id
					this.field.Classshape_FieldsDBID.Valid = true
					this.field.Classshape_FieldsDBID_Index = new NullInt64
					this.field.Classshape_FieldsDBID_Index.Valid = true
					break
			}
			this.fieldService.postField(this.field).subscribe(field => {

				this.fieldService.FieldServiceChanged.next("post")

				this.field = {} // reset fields
			});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.width = "50%"
		dialogConfig.height = "50%"
		dialogConfig.data = {
			ID: this.field.ID,
			ReversePointer: reverseField,
			OrderingMode: false,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'sTableComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
		});
	}

	openDragAndDropOrdering(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.data = {
			ID: this.field.ID,
			ReversePointer: reverseField,
			OrderingMode: true,
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

	fillUpNameIfEmpty(event) {
		if (this.field.Name == undefined) {
			this.field.Name = event.value.Name		
		}
	}
}
