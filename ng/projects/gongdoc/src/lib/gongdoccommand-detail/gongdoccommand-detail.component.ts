// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { GongdocCommandDB } from '../gongdoccommand-db'
import { GongdocCommandService } from '../gongdoccommand.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { GongdocCommandTypeSelect, GongdocCommandTypeList } from '../GongdocCommandType'
import { GongdocNodeTypeSelect, GongdocNodeTypeList } from '../GongdocNodeType'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-gongdoccommand-detail',
	templateUrl: './gongdoccommand-detail.component.html',
	styleUrls: ['./gongdoccommand-detail.component.css'],
})
export class GongdocCommandDetailComponent implements OnInit {

	// insertion point for declarations
	GongdocCommandTypeList: GongdocCommandTypeSelect[]
	GongdocNodeTypeList: GongdocNodeTypeSelect[]

	// the GongdocCommandDB of interest
	gongdoccommand: GongdocCommandDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private gongdoccommandService: GongdocCommandService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {
		this.getGongdocCommand()

		// observable for changes in structs
		this.gongdoccommandService.GongdocCommandServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getGongdocCommand()
				}
			}
		)

		// insertion point for initialisation of enums list
		this.GongdocCommandTypeList = GongdocCommandTypeList
		this.GongdocNodeTypeList = GongdocNodeTypeList
	}

	getGongdocCommand(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				if (id != 0 && association == undefined) {
					this.gongdoccommand = frontRepo.GongdocCommands.get(id)
				} else {
					this.gongdoccommand = new (GongdocCommandDB)
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
		}

		if (id != 0 && association == undefined) {

			this.gongdoccommandService.updateGongdocCommand(this.gongdoccommand)
				.subscribe(gongdoccommand => {
					this.gongdoccommandService.GongdocCommandServiceChanged.next("update")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
			}
			this.gongdoccommandService.postGongdocCommand(this.gongdoccommand).subscribe(gongdoccommand => {

				this.gongdoccommandService.GongdocCommandServiceChanged.next("post")

				this.gongdoccommand = {} // reset fields
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
			ID: this.gongdoccommand.ID,
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
			ID: this.gongdoccommand.ID,
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
		if (this.gongdoccommand.Name == undefined) {
			this.gongdoccommand.Name = event.value.Name		
		}
	}
}
