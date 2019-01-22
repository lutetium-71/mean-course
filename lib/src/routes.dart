import 'package:angular_router/angular_router.dart';

import 'posts/post_list/post_list_component.template.dart'
    as post_list_template;
import 'posts/post_create/post_create_component.template.dart'
    as post_create_template;
import 'posts/post_edit/post_edit_component.template.dart'
    as post_edit_template;
import 'auth/login/login_component.template.dart' as login_template;
import 'auth/signup/signup_component.template.dart' as signup_template;

import 'route_paths.dart';
export 'route_paths.dart';

class Routes {
  static final posts = RouteDefinition(
    routePath: RoutePaths.posts,
    component: post_list_template.PostListComponentNgFactory,
  );
  static final create = RouteDefinition(
    routePath: RoutePaths.create,
    component: post_create_template.PostCreateComponentNgFactory,
  );
  static final edit = RouteDefinition(
    routePath: RoutePaths.edit,
    component: post_edit_template.PostEditComponentNgFactory,
  );
  static final login = RouteDefinition(
    routePath: RoutePaths.login,
    component: login_template.LoginComponentNgFactory,
  );
  static final signup = RouteDefinition(
    routePath: RoutePaths.signup,
    component: signup_template.SignUpComponentNgFactory,
  );

  static final all = <RouteDefinition>[
    posts,
    create,
    edit,
    login,
    signup,
  ];
}
