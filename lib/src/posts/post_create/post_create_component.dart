import 'package:angular/angular.dart';
import 'package:angular_forms/angular_forms.dart';
import 'package:angular_components/material_button/material_button.dart';
import 'package:angular_components/material_icon/material_icon.dart';
import 'package:angular_components/material_input/material_input.dart';

@Component(
    selector: 'app-post-create',
    styleUrls: [
      'package:angular_components/css/mdc_web/card/mdc-card.scss.css',
      'post_create_component.css'
    ],
    templateUrl: 'post_create_component.html',
    directives: [
      formDirectives,
      MaterialButtonComponent,
      MaterialIconComponent,
      materialInputDirectives,
    ])
class PostCreateComponent {
  var enteredValue = '';
  var newPost = 'NO CONTENT';

  onAddPost() {
    this.newPost = this.enteredValue;
  }
}
