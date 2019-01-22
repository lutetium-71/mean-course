import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';
import 'package:angular_router/angular_router.dart';

import 'src/posts/post_create/post_create_component.dart';
import 'src/posts/post_list/post_list_component.dart';
import 'src/header/header_component.dart';
import 'src/auth/signup/signup_component.dart';
import 'src/auth/login/login_component.dart';
import 'src/post_list_service.dart';
import 'src/auth_service.dart';
import 'src/routes.dart';

@Component(
  selector: 'my-app',
  styleUrls: ['app_component.css'],
  templateUrl: 'app_component.html',
  directives: [
    HeaderComponent,
    PostCreateComponent,
    PostListComponent,
    LoginComponent,
    SignUpComponent,
    routerDirectives,
  ],
  providers: [
    materialProviders,
    ClassProvider(PostListService),
    ClassProvider(AuthService)
  ],
  exports: [RoutePaths, Routes],
)
class AppComponent {}
