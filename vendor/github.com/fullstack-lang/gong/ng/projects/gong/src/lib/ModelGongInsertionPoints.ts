// generated from ng_file_enum.ts.go
export enum ModelGongInsertionPoints {
	// insertion point	
	ModelGongInsertionArrayDefintion = 4,
	ModelGongInsertionArrayInitialisation = 5,
	ModelGongInsertionArrayNil = 7,
	ModelGongInsertionArrayReset = 6,
	ModelGongInsertionCommitCheckoutSignature = 0,
	ModelGongInsertionCreateCallback = 2,
	ModelGongInsertionDeleteCallback = 3,
	ModelGongInsertionStageFunctions = 1,
	ModelGongInsertionsNb = 8,
}

export interface ModelGongInsertionPointsSelect {
	value: string;
	viewValue: string;
}

export const ModelGongInsertionPointsList: ModelGongInsertionPointsSelect[] = [ // insertion point	
	{ value: 'ModelGongInsertionArrayDefintion', viewValue: '4' },
	{ value: 'ModelGongInsertionArrayInitialisation', viewValue: '5' },
	{ value: 'ModelGongInsertionArrayNil', viewValue: '7' },
	{ value: 'ModelGongInsertionArrayReset', viewValue: '6' },
	{ value: 'ModelGongInsertionCommitCheckoutSignature', viewValue: '0' },
	{ value: 'ModelGongInsertionCreateCallback', viewValue: '2' },
	{ value: 'ModelGongInsertionDeleteCallback', viewValue: '3' },
	{ value: 'ModelGongInsertionStageFunctions', viewValue: '1' },
	{ value: 'ModelGongInsertionsNb', viewValue: '8' },
];