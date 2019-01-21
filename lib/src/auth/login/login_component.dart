import 'package:angular/angular.dart';
import 'package:angular_router/angular_router.dart';
import 'package:angular_forms/angular_forms.dart';
import 'package:angular_components/focus/focus.dart';
import 'package:angular_components/material_button/material_button.dart';
import 'package:angular_components/material_input/material_input.dart';

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
  onLogin(NgForm form) {
    print(form.value);
  }
}
