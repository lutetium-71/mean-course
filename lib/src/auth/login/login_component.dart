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
    'login_component.css'
  ],
  templateUrl: 'login_component.html',
  directives: [
    AutoFocusDirective,
    routerDirectives,
    formDirectives,
    MaterialButtonComponent,
    materialInputDirectives,
    NgForm,
  ],
)
class LoginComponent {
  AuthService _authService;

  LoginComponent(this._authService);

  onLogin(NgForm form) {
    User post = User(null, form.value["email"], form.value["password"]);
    _authService.loginUser(post);
    // form.reset();
  }
}
