import 'package:angular_router/angular_router.dart';

import './posts/post_list/post_list_component.template.dart'
    as post_list_template;
import './posts/post_create/post_create_component.template.dart'
    as post_create_template;

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

  static final all = <RouteDefinition>[
    posts,
    create,
  ];
}
