import {Component, ElementRef, OnInit, ViewChild} from '@angular/core';
import {
  trigger,
  state,
  style,
  animate,
  transition
} from '@angular/animations';


const TEXT_ANIMATION__COMMON = [
  state('true', style({transform: 'translateX(0)', visibility: 'visible'})),
  state('false', style({transform: 'translateX(-10vw)', visibility: 'hidden'})),
  transition('false => true', animate('0.4s ease-in'))
];

const MUSICAL_NOTE_ANIMATION = [
  state('true', style({transform: 'translateX(0)', visibility: 'visible'})),
  state('false', style({transform: 'translateX(10vw)', visibility: 'hidden'})),
  transition('false => true', animate('0.4s ease-in'))
];


@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.scss'],
  animations: [
    trigger('musicalNoteVisibleState', MUSICAL_NOTE_ANIMATION),
    trigger('searchVisibleState', TEXT_ANIMATION__COMMON),
    trigger('listenVisibleState', TEXT_ANIMATION__COMMON),
    trigger('uploadVisibleState', TEXT_ANIMATION__COMMON),
  ],
  host: {
    '(window:scroll)': 'updateElementsBasedOnScrollEvent($event)'
  }
})
export class HomepageComponent implements OnInit {
  @ViewChild('musicalNote') musicalNoteEle;
  @ViewChild('searchText') searchTextEle;
  @ViewChild('listenText') listenTextEle;
  @ViewChild('uploadText') uploadTextEle;

  elementDict: any;

  constructor() {
  }

  ngOnInit() {
    this.elementDict = {
      musicalNote: {
        viewChild: this.musicalNoteEle,
        isVisible: "false"
      },
      searchText: {
        viewChild: this.searchTextEle,
        isVisible: "false"
      },
      listenText: {
        viewChild: this.listenTextEle,
        isVisible: "false"
      },
      uploadText: {
        viewChild: this.uploadTextEle,
        isVisible: "false"
      }
    }
  }

  /**
   This function is called When the user scrolls around. If the user scrolls
   into any element listed in 'this.elementDict,' that element becomes visible
   and slides-in into view. Once an element becomes visible, it'll be categorized
   as 'always-visible.'
   **/
  updateElementsBasedOnScrollEvent(evt) {
    for (let eleKey in this.elementDict) {
      let ele = this.elementDict[eleKey];
      if (this.elementDict.hasOwnProperty(eleKey) && ele.isVisible != "true") {
        let rect = ele.viewChild.nativeElement.getBoundingClientRect();
        ele.isVisible = this.isVisible(rect.bottom, window.innerHeight || document.documentElement.clientHeight);
      }
    }
  }

  /**
   Helper function determines if the bottom point of an element is visible
   within the window's height.
   **/
  isVisible(elementBottomPosition: number, windowHeight: number): string {
    return String((elementBottomPosition - windowHeight) < 0);
  }
}
