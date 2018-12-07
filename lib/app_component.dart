import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';

import 'src/posts/post_create/post_create_component.dart';
import 'src/posts/post_list/post_list_component.dart';
import 'src/header/header_component.dart';

@Component(
  selector: 'my-app',
  styleUrls: ['app_component.css'],
  templateUrl: 'app_component.html',
  directives: [PostCreateComponent, PostListComponent, HeaderComponent],
  providers: const <dynamic>[materialProviders],
)
class AppComponent {
  // Nothing here yet. All logic is in TodoListComponent.
}
