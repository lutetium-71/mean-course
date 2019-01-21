import 'package:angular/angular.dart';
import 'package:angular_router/angular_router.dart';
import 'package:angular_components/material_button/material_button.dart';

@Component(
  selector: 'app-signup',
  styleUrls: [
    'package:angular_components/css/mdc_web/card/mdc-card.scss.css',
    'signup_component.css'
  ],
  templateUrl: 'signup_component.html',
  directives: [routerDirectives, MaterialButtonComponent],
)
class SignUpComponent {}
