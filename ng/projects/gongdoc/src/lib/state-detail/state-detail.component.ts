// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { StateDB } from '../state-db'
import { StateService } from '../state.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-state-detail',
	templateUrl: './state-detail.component.html',
	styleUrls: ['./state-detail.component.css'],
})
export class StateDetailComponent implements OnInit {

	// insertion point for declarations

	// the StateDB of interest
	state: StateDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private stateService: StateService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {
		this.getState()

		// observable for changes in structs
		this.stateService.StateServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getState()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getState(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				if (id != 0 && association == undefined) {
					this.state = frontRepo.States.get(id)
				} else {
					this.state = new (StateDB)
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
			if (this.state.Umlsc_States_reverse != undefined) {
				this.state.Umlsc_StatesDBID = new NullInt64
				this.state.Umlsc_StatesDBID.Int64 = this.state.Umlsc_States_reverse.ID
				this.state.Umlsc_StatesDBID.Valid = true
				this.state.Umlsc_StatesDBID_Index = new NullInt64
				this.state.Umlsc_StatesDBID_Index.Valid = true
				this.state.Umlsc_States_reverse = undefined // very important, otherwise, circular JSON
			}
		}

		if (id != 0 && association == undefined) {

			this.stateService.updateState(this.state)
				.subscribe(state => {
					this.stateService.StateServiceChanged.next("update")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "Umlsc_States":
					this.state.Umlsc_StatesDBID = new NullInt64
					this.state.Umlsc_StatesDBID.Int64 = id
					this.state.Umlsc_StatesDBID.Valid = true
					this.state.Umlsc_StatesDBID_Index = new NullInt64
					this.state.Umlsc_StatesDBID_Index.Valid = true
					break
			}
			this.stateService.postState(this.state).subscribe(state => {

				this.stateService.StateServiceChanged.next("post")

				this.state = {} // reset fields
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
			ID: this.state.ID,
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
			ID: this.state.ID,
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
		if (this.state.Name == undefined) {
			this.state.Name = event.value.Name		
		}
	}
}
