import { Component, OnInit } from '@angular/core';
import { Observable, timer } from 'rxjs';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import * as gongdoc from 'gongdoc'

/**
 * Food data with nested structure.
 * Each node has a name and an optional list of children.
 */
interface Node {
  name: string;

  gongNode: gongdoc.NodeDB;
  children?: Node[];
}

/** Flat node with expandable and level information */
interface FlatNode {
  expandable: boolean;

  name: string;
  gongNode: gongdoc.NodeDB;
  level: number;
}

@Component({
  selector: 'lib-tree',
  templateUrl: './tree.component.html',
  styleUrls: ['./tree.component.css']
})
export class TreeComponent implements OnInit {

  private _transformer = (node: Node, level: number) => {
    return {
      expandable: !!node.children && node.children.length > 0,

      gongNode: node.gongNode,
      name: node.name,
      level: level,
    };
  };

  treeControl = new FlatTreeControl<FlatNode>(
    node => node.level,
    node => node.expandable,
  );

  treeFlattener = new MatTreeFlattener(
    this._transformer,
    node => node.level,
    node => node.expandable,
    node => node.children,
  );

  dataSource = new MatTreeFlatDataSource(this.treeControl, this.treeFlattener);

  hasChild = (_: number, node: FlatNode) => node.expandable;

  public gongdocFrontRepo?: gongdoc.FrontRepo

  constructor(
    private gongdocFrontRepoService: gongdoc.FrontRepoService,
    private gongdocCommitNbService: gongdoc.CommitNbService,
    private gongdocPushFromFrontNbService: gongdoc.PushFromFrontNbService,
    private gongdocNodeService: gongdoc.NodeService,

  ) {
    // this.dataSource.data = TREE_DATA;
  }


  // the component is refreshed when modification are performed in the back repo 
  // 
  // the checkCommitNbTimer polls the commit number of the back repo
  // if the commit number has increased, it pulls the front repo and redraw the diagram

  checkCommitNbTimer: Observable<number> = timer(500, 500);
  lastCommitNb = -1
  lastPushFromFrontNb = -1
  currTime: number = 0

  ngOnInit(): void {

    this.checkCommitNbTimer.subscribe(
      currTime => {
        this.currTime = currTime

        // see above for the explanation
        this.gongdocCommitNbService.getCommitNb().subscribe(
          commitNb => {
            if (this.lastCommitNb < commitNb) {

              console.log("last commit nb " + this.lastCommitNb + " new: " + commitNb)
              this.refresh()
              this.lastCommitNb = commitNb
            }
          }
        )

        // see above for the explanation
        this.gongdocPushFromFrontNbService.getPushFromFrontNb().subscribe(
          pushFromFrontNb => {
            if (this.lastPushFromFrontNb < pushFromFrontNb) {

              console.log("last commit nb " + this.lastPushFromFrontNb + " new: " + pushFromFrontNb)
              this.refresh()
              this.lastPushFromFrontNb = pushFromFrontNb
            }
          }
        )
      }
    )
  }

  refresh(): void {

    this.gongdocFrontRepoService.pull().subscribe(
      gongdocsFrontRepo => {
        this.gongdocFrontRepo = gongdocsFrontRepo


        if (this.gongdocFrontRepo.Trees_array.length != 1) {
          console.log("error: there should be exactly one tree")
          return
        }
        var treeSingloton = this.gongdocFrontRepo.Trees_array[0]

        var rootNodes = new Array<Node>()

        if (treeSingloton.RootNodes != undefined) {
          rootNodes = treeSingloton.RootNodes.map(nodeDB => this.gongNodeToMatTreeNode(nodeDB))
        }

        this.dataSource.data = rootNodes

        // expand nodes that were exapanded before
        this.treeControl.dataNodes?.forEach(
          node => {
            if (node.gongNode.IsExpanded) {
              this.treeControl.expand(node)
            }
          }
        )

      }
    )
  }

  gongNodeToMatTreeNode(nodeDB: gongdoc.NodeDB): Node {
    var matTreeNode: Node = { name: nodeDB.Name, gongNode: nodeDB, children: [] }
    if (nodeDB.Children != undefined) {
      matTreeNode.children = nodeDB.Children.map(child => this.gongNodeToMatTreeNode(child))
    }

    return matTreeNode
  }

  toggleNodeExpansion(node: FlatNode): void {
    console.log(node.name)

    node.gongNode.IsExpanded = !node.gongNode.IsExpanded

    this.gongdocNodeService.updateNode(node.gongNode).subscribe(
      gongdocNode => {
        console.log("updated node")
      }
    )
  }

  toggleNodeCheckbox(node: FlatNode): void {
    console.log(node.name)

    node.gongNode.IsChecked = !node.gongNode.IsChecked

    this.gongdocNodeService.updateNode(node.gongNode).subscribe(
      gongdocNode => {
        console.log("updated node")
      }
    )
  }
}
