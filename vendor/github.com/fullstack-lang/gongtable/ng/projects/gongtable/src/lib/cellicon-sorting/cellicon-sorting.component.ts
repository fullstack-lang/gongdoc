// generated by gong
import { Component, OnInit, Inject, Optional } from '@angular/core';
import { TypeofExpr } from '@angular/compiler';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { CellIconDB } from '../cellicon-db'
import { CellIconService } from '../cellicon.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { NullInt64 } from '../null-int64'

@Component({
  selector: 'lib-cellicon-sorting',
  templateUrl: './cellicon-sorting.component.html',
  styleUrls: ['./cellicon-sorting.component.css']
})
export class CellIconSortingComponent implements OnInit {

  frontRepo: FrontRepo = new (FrontRepo)

  // array of CellIcon instances that are in the association
  associatedCellIcons = new Array<CellIconDB>();

  constructor(
    private celliconService: CellIconService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of cellicon instances
    public dialogRef: MatDialogRef<CellIconSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getCellIcons()
  }

  getCellIcons(): void {
    this.frontRepoService.pull(this.dialogData.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let index = 0
        for (let cellicon of this.frontRepo.CellIcons_array) {
          let ID = this.dialogData.ID
          let revPointerID = cellicon[this.dialogData.ReversePointer as keyof CellIconDB] as unknown as NullInt64
          let revPointerID_Index = cellicon[this.dialogData.ReversePointer + "_Index" as keyof CellIconDB] as unknown as NullInt64
          if (revPointerID.Int64 == ID) {
            if (revPointerID_Index == undefined) {
              revPointerID_Index = new NullInt64
              revPointerID_Index.Valid = true
              revPointerID_Index.Int64 = index++
            }
            this.associatedCellIcons.push(cellicon)
          }
        }

        // sort associated cellicon according to order
        this.associatedCellIcons.sort((t1, t2) => {
          let t1_revPointerID_Index = t1[this.dialogData.ReversePointer + "_Index" as keyof typeof t1] as unknown as NullInt64
          let t2_revPointerID_Index = t2[this.dialogData.ReversePointer + "_Index" as keyof typeof t2] as unknown as NullInt64
          if (t1_revPointerID_Index && t2_revPointerID_Index) {
            if (t1_revPointerID_Index.Int64 > t2_revPointerID_Index.Int64) {
              return 1;
            }
            if (t1_revPointerID_Index.Int64 < t2_revPointerID_Index.Int64) {
              return -1;
            }
          }
          return 0;
        });
      }
    )
  }

  drop(event: CdkDragDrop<string[]>) {
    moveItemInArray(this.associatedCellIcons, event.previousIndex, event.currentIndex);

    // set the order of CellIcon instances
    let index = 0

    for (let cellicon of this.associatedCellIcons) {
      let revPointerID_Index = cellicon[this.dialogData.ReversePointer + "_Index" as keyof CellIconDB] as unknown as NullInt64
      revPointerID_Index.Valid = true
      revPointerID_Index.Int64 = index++
    }
  }

  save() {

    this.associatedCellIcons.forEach(
      cellicon => {
        this.celliconService.updateCellIcon(cellicon, this.dialogData.GONG__StackPath)
          .subscribe(cellicon => {
            this.celliconService.CellIconServiceChanged.next("update")
          });
      }
    )

    this.dialogRef.close('Sorting of ' + this.dialogData.ReversePointer + ' done');
  }
}