import 'package:angular/angular.dart';

import 'src/posts/post_create/post_create_component.dart';

// AngularDart info: https://webdev.dartlang.org/angular
// Components info: https://webdev.dartlang.org/components

@Component(
  selector: 'my-app',
  styleUrls: ['app_component.css'],
  templateUrl: 'app_component.html',
  directives: [PostCreateComponent],
)
class AppComponent {
  // Nothing here yet. All logic is in TodoListComponent.
}
