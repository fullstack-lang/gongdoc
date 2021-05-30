// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { ClassdiagramDB } from '../classdiagram-db'
import { ClassdiagramService } from '../classdiagram.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-classdiagram-detail',
	templateUrl: './classdiagram-detail.component.html',
	styleUrls: ['./classdiagram-detail.component.css'],
})
export class ClassdiagramDetailComponent implements OnInit {

	// insertion point for declarations

	// the ClassdiagramDB of interest
	classdiagram: ClassdiagramDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private classdiagramService: ClassdiagramService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {
		this.getClassdiagram()

		// observable for changes in structs
		this.classdiagramService.ClassdiagramServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getClassdiagram()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getClassdiagram(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				if (id != 0 && association == undefined) {
					this.classdiagram = frontRepo.Classdiagrams.get(id)
				} else {
					this.classdiagram = new (ClassdiagramDB)
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
			if (this.classdiagram.Pkgelt_Classdiagrams_reverse != undefined) {
				this.classdiagram.Pkgelt_ClassdiagramsDBID = new NullInt64
				this.classdiagram.Pkgelt_ClassdiagramsDBID.Int64 = this.classdiagram.Pkgelt_Classdiagrams_reverse.ID
				this.classdiagram.Pkgelt_ClassdiagramsDBID.Valid = true
				this.classdiagram.Pkgelt_ClassdiagramsDBID_Index = new NullInt64
				this.classdiagram.Pkgelt_ClassdiagramsDBID_Index.Valid = true
				this.classdiagram.Pkgelt_Classdiagrams_reverse = undefined // very important, otherwise, circular JSON
			}
		}

		if (id != 0 && association == undefined) {

			this.classdiagramService.updateClassdiagram(this.classdiagram)
				.subscribe(classdiagram => {
					this.classdiagramService.ClassdiagramServiceChanged.next("update")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "Pkgelt_Classdiagrams":
					this.classdiagram.Pkgelt_ClassdiagramsDBID = new NullInt64
					this.classdiagram.Pkgelt_ClassdiagramsDBID.Int64 = id
					this.classdiagram.Pkgelt_ClassdiagramsDBID.Valid = true
					this.classdiagram.Pkgelt_ClassdiagramsDBID_Index = new NullInt64
					this.classdiagram.Pkgelt_ClassdiagramsDBID_Index.Valid = true
					break
			}
			this.classdiagramService.postClassdiagram(this.classdiagram).subscribe(classdiagram => {

				this.classdiagramService.ClassdiagramServiceChanged.next("post")

				this.classdiagram = {} // reset fields
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
			ID: this.classdiagram.ID,
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
			ID: this.classdiagram.ID,
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
}
