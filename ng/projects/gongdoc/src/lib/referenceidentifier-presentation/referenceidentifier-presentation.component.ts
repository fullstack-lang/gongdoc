import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { ReferenceIdentifierDB } from '../referenceidentifier-db'
import { ReferenceIdentifierService } from '../referenceidentifier.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface referenceidentifierDummyElement {
}

const ELEMENT_DATA: referenceidentifierDummyElement[] = [
];

@Component({
	selector: 'app-referenceidentifier-presentation',
	templateUrl: './referenceidentifier-presentation.component.html',
	styleUrls: ['./referenceidentifier-presentation.component.css'],
})
export class ReferenceIdentifierPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	referenceidentifier: ReferenceIdentifierDB = new (ReferenceIdentifierDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private referenceidentifierService: ReferenceIdentifierService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getReferenceIdentifier();

		// observable for changes in 
		this.referenceidentifierService.ReferenceIdentifierServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getReferenceIdentifier()
				}
			}
		)
	}

	getReferenceIdentifier(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.referenceidentifier = this.frontRepo.ReferenceIdentifiers.get(id)!

				// insertion point for recovery of durations
				// insertion point for recovery of enum tint
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gongdoc_go_presentation: ["github_com_fullstack_lang_gongdoc_go-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gongdoc_go_editor: ["github_com_fullstack_lang_gongdoc_go-" + "referenceidentifier-detail", ID]
			}
		}]);
	}
}
