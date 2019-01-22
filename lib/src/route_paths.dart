import 'package:angular_router/angular_router.dart';

const postId = 'id';

class RoutePaths {
  static final posts = RoutePath(path: '');
  static final create = RoutePath(path: 'create');
  static final edit = RoutePath(path: 'edit/:$postId');
  static final login = RoutePath(path: 'login');
  static final signup = RoutePath(path: 'signup');
}

String getId(Map<String, String> parameters) {
  final id = parameters[postId];
  return id == null ? null : id;
}
