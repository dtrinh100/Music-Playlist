import {Directive, Input, HostListener} from '@angular/core';


@Directive({
    selector: '[validateOnBlur]',
    })

  export class RegistrationDirective {
    @Input('validateFormControl') validateFormControl;

    constructor() { }
    @HostListener('focus', ['$event.target'])
      onFocus(target) {
        this.validateFormControl.markAsUntouched();

    }
    @HostListener('focusout', ['$event.target'])
    onFocusout(target) {
      console.log("Focus out called");
      this.validateFormControl.markAsTouched();
    }
  }
