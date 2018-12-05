import 'package:angular/angular.dart';
import 'package:angular_forms/angular_forms.dart';

@Component(
    selector: 'app-post-create',
    styleUrls: ['post_create_component.css'],
    templateUrl: 'post_create_component.html',
    directives: [formDirectives])
class PostCreateComponent {
  var enteredValue = '';
  var newPost = 'NO CONTENT';

  onAddPost() {
    this.newPost = this.enteredValue;
  }
}
