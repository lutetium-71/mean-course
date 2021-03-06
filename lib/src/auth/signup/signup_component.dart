import 'package:angular/angular.dart';
import 'package:angular_router/angular_router.dart';
import 'package:angular_forms/angular_forms.dart';
import 'package:angular_components/focus/focus.dart';
import 'package:angular_components/material_button/material_button.dart';
import 'package:angular_components/material_input/material_input.dart';

import '../../user_model.dart';
import '../../auth_service.dart';

@Component(
  selector: 'app-signup',
  styleUrls: [
    'package:angular_components/css/mdc_web/card/mdc-card.scss.css',
    'signup_component.css'
  ],
  templateUrl: 'signup_component.html',
  directives: [
    AutoFocusDirective,
    routerDirectives,
    formDirectives,
    MaterialButtonComponent,
    materialInputDirectives,
    NgForm,
  ],
)
class SignUpComponent {
  AuthService _authService;

  SignUpComponent(this._authService);

  onSignUp(NgForm form) {
    User post = User(null, form.value["email"], form.value["password"]);
    _authService.createUser(post);
    // form.reset();
  }
}
