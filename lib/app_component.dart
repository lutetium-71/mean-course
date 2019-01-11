import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';

import 'src/posts/post_create/post_create_component.dart';
import 'src/posts/post_list/post_list_component.dart';
import 'src/header/header_component.dart';
import 'src/post_list_service.dart';

@Component(
  selector: 'my-app',
  styleUrls: ['app_component.css'],
  templateUrl: 'app_component.html',
  directives: [PostCreateComponent, PostListComponent, HeaderComponent],
  providers: [materialProviders, ClassProvider(PostListService)],
)
class AppComponent {}
