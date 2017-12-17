import {
  Directive,
  Input,
  OnInit,
  TemplateRef,
  ViewContainerRef
} from '@angular/core';

import { AuthService } from '../services';

@Directive({selector: '[appShowAuthed]'})
export class ShowAuthedDirective implements OnInit {
  condition: boolean;

  @Input() set appShowAuthed(condition: boolean) {
    this.condition = condition;
  }

  constructor(private templateRef: TemplateRef<any>,
              private authService: AuthService,
              private viewContainer: ViewContainerRef) {
  }

  ngOnInit(): void {
    this.authService.isAuthenticated.subscribe(
      (isAuthenticated) => {
        this.viewContainer.clear();
        if ((isAuthenticated && this.condition) || (!isAuthenticated && !this.condition)) {
          this.viewContainer.createEmbeddedView(this.templateRef);
        }
      }
    )
  }
}
