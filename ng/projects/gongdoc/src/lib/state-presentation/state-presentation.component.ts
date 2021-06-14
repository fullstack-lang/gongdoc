import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { StateDB } from '../state-db'
import { StateService } from '../state.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface stateDummyElement {
}

const ELEMENT_DATA: stateDummyElement[] = [
];

@Component({
	selector: 'app-state-presentation',
	templateUrl: './state-presentation.component.html',
	styleUrls: ['./state-presentation.component.css'],
})
export class StatePresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	state: StateDB;

	// front repo
	frontRepo: FrontRepo
 
	constructor(
		private stateService: StateService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getState();

		// observable for changes in 
		this.stateService.StateServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getState()
				}
			}
		)
	}

	getState(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.state = this.frontRepo.States.get(id)

				// insertion point for recovery of durations
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
				github_com_fullstack_lang_gongdoc_go_editor: ["github_com_fullstack_lang_gongdoc_go-" + "state-detail", ID]
			}
		}]);
	}
}
