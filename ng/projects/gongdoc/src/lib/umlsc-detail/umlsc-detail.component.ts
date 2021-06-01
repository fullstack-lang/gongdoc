// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { UmlscDB } from '../umlsc-db'
import { UmlscService } from '../umlsc.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-umlsc-detail',
	templateUrl: './umlsc-detail.component.html',
	styleUrls: ['./umlsc-detail.component.css'],
})
export class UmlscDetailComponent implements OnInit {

	// insertion point for declarations

	// the UmlscDB of interest
	umlsc: UmlscDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private umlscService: UmlscService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {
		this.getUmlsc()

		// observable for changes in structs
		this.umlscService.UmlscServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getUmlsc()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getUmlsc(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				if (id != 0 && association == undefined) {
					this.umlsc = frontRepo.Umlscs.get(id)
				} else {
					this.umlsc = new (UmlscDB)
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
			if (this.umlsc.Pkgelt_Umlscs_reverse != undefined) {
				this.umlsc.Pkgelt_UmlscsDBID = new NullInt64
				this.umlsc.Pkgelt_UmlscsDBID.Int64 = this.umlsc.Pkgelt_Umlscs_reverse.ID
				this.umlsc.Pkgelt_UmlscsDBID.Valid = true
				this.umlsc.Pkgelt_UmlscsDBID_Index = new NullInt64
				this.umlsc.Pkgelt_UmlscsDBID_Index.Valid = true
				this.umlsc.Pkgelt_Umlscs_reverse = undefined // very important, otherwise, circular JSON
			}
		}

		if (id != 0 && association == undefined) {

			this.umlscService.updateUmlsc(this.umlsc)
				.subscribe(umlsc => {
					this.umlscService.UmlscServiceChanged.next("update")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "Pkgelt_Umlscs":
					this.umlsc.Pkgelt_UmlscsDBID = new NullInt64
					this.umlsc.Pkgelt_UmlscsDBID.Int64 = id
					this.umlsc.Pkgelt_UmlscsDBID.Valid = true
					this.umlsc.Pkgelt_UmlscsDBID_Index = new NullInt64
					this.umlsc.Pkgelt_UmlscsDBID_Index.Valid = true
					break
			}
			this.umlscService.postUmlsc(this.umlsc).subscribe(umlsc => {

				this.umlscService.UmlscServiceChanged.next("post")

				this.umlsc = {} // reset fields
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
			ID: this.umlsc.ID,
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
			ID: this.umlsc.ID,
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
		if (this.umlsc.Name == undefined) {
			this.umlsc.Name = event.value.Name		
		}
	}
}
